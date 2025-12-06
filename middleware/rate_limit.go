package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"PennieAI/config"
)

type RateLimitConfig struct {
	MaxRequests int           // Maximum number of requests allowed
	Window      time.Duration // Time window for the limit (e.g., 1 hour)
	KeyPrefix   string        // Prefix for Redis keys (e.g., "ratelimit:openai:")
}

var DefaultOpenAIRateLimit = RateLimitConfig{
	MaxRequests: 100,                // 100 requests per window
	Window:      time.Hour,          // 1 hour window
	KeyPrefix:   "ratelimit:openai", // Redis key prefix
}

// RateLimiter creates a rate limiting middleware using Redis
func RateLimiter(rateLimitConfig RateLimitConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		rdb := config.GetRedis()

		userID := "::123"

		key := fmt.Sprintf("%s:%s", rateLimitConfig.KeyPrefix, userID)

		// Create context with timeout to prevent hanging
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		// Try to increment the counter
		allowed, remaining, resetTime, err := checkRateLimit(ctx, rdb, key, rateLimitConfig)

		if err != nil {
			// If Redis fails, we log the error but allow the request through
			// This is called "fail open" - we prefer availability over strict rate limiting
			fmt.Printf("⚠️  Rate limit check failed: %v\n", err)
			c.Next()
			return
		}

		// Add rate limit headers to response (standard practice)
		c.Header("X-RateLimit-Limit", strconv.Itoa(rateLimitConfig.MaxRequests))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(remaining))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(resetTime.Unix(), 10))

		if !allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":       "Rate limit exceeded",
				"message":     fmt.Sprintf("You have exceeded the rate limit of %d requests per %v", rateLimitConfig.MaxRequests, rateLimitConfig.Window),
				"retry_after": resetTime.Sub(time.Now()).Seconds(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// checkRateLimit implements the rate limiting algorithm using Redis
func checkRateLimit(ctx context.Context, rdb *redis.Client, key string, rateLimitConfig RateLimitConfig) (bool, int, time.Time, error) {
	// Use Redis INCR to atomically increment the counter
	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return false, 0, time.Time{}, err
	}

	// If this is the first request (count == 1), set the expiration
	if count == 1 {
		err = rdb.Expire(ctx, key, rateLimitConfig.Window).Err()
		if err != nil {
			return false, 0, time.Time{}, err
		}
	}

	// Get the TTL (time to live) to know when the limit resets
	ttl, err := rdb.TTL(ctx, key).Result()
	if err != nil {
		return false, 0, time.Time{}, err
	}

	resetTime := time.Now().Add(ttl)

	remaining := rateLimitConfig.MaxRequests - int(count)
	if remaining < 0 {
		remaining = 0
	}

	allowed := count <= int64(rateLimitConfig.MaxRequests)

	return allowed, remaining, resetTime, nil
}

func OpenAIRateLimiter() gin.HandlerFunc {
	return RateLimiter(DefaultOpenAIRateLimit)
}
