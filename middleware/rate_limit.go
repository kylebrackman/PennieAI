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
	MaxRequests int
	Window      time.Duration
	KeyPrefix   string
}

type RateLimitResult struct {
	Allowed   bool
	Remaining int
	ResetTime time.Time
}

var DefaultOpenAIRateLimit = RateLimitConfig{
	MaxRequests: 100,
	Window:      time.Hour,
	KeyPrefix:   "ratelimit:openai", // Redis key prefix
}

func RateLimiter(rateLimitConfig RateLimitConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		rdb := config.GetRedis()

		userID := "::123" // Placeholder

		key := fmt.Sprintf("%s:%s", rateLimitConfig.KeyPrefix, userID)

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		erlResult, err := enforceRateLimit(ctx, rdb, key, rateLimitConfig)

		if err != nil {
			// If Redis fails, log the error but allow the request through
			// Fail open - prefer availability over strict rate limiting
			fmt.Printf("⚠️  Rate limit check failed: %v\n", err)

			// Move to the next handler in the chain (skip rate limiting)
			c.Next()
			return
		}

		// Add rate limit headers to response (standard practice)
		c.Header("X-RateLimit-Limit", strconv.Itoa(rateLimitConfig.MaxRequests))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(erlResult.Remaining))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(erlResult.ResetTime.Unix(), 10))

		if !erlResult.Allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":       "Rate limit exceeded",
				"message":     fmt.Sprintf("You have exceeded the rate limit of %d requests per %v", rateLimitConfig.MaxRequests, rateLimitConfig.Window),
				"retry_after": erlResult.ResetTime.Sub(time.Now()).Seconds(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func enforceRateLimit(ctx context.Context, rdb *redis.Client, key string, rateLimitConfig RateLimitConfig) (*RateLimitResult, error) {
	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if count == 1 {
		err = rdb.Expire(ctx, key, rateLimitConfig.Window).Err()
		if err != nil {
			return nil, err
		}
	}

	ttl, err := rdb.TTL(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	resetTime := time.Now().Add(ttl)

	remaining := rateLimitConfig.MaxRequests - int(count)
	if remaining < 0 {
		remaining = 0
	}

	allowed := count <= int64(rateLimitConfig.MaxRequests)

	return &RateLimitResult{
		Allowed:   allowed,
		Remaining: remaining,
		ResetTime: resetTime,
	}, nil
}

func OpenAIRateLimiter() gin.HandlerFunc {
	return RateLimiter(DefaultOpenAIRateLimit)
}
