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

		allowed, remaining, resetTime, err := enforceRateLimit(ctx, rdb, key, rateLimitConfig)

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

func enforceRateLimit(ctx context.Context, rdb *redis.Client, key string, rateLimitConfig RateLimitConfig) (bool, int, time.Time, error) {
	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return false, 0, time.Time{}, err
	}

	if count == 1 {
		err = rdb.Expire(ctx, key, rateLimitConfig.Window).Err() // What is .Err()?
		if err != nil {
			return false, 0, time.Time{}, err
		}
	}

	// Get the TTL (time to live) to know when the limit resets
	ttl, err := rdb.TTL(ctx, key).Result() // why do we pass ctx again here? is it so that this entire flow can be cancelled if it's been over 2 seconds overall?
	if err != nil {
		return false, 0, time.Time{}, err
	}

	resetTime := time.Now().Add(ttl) // Why are we adding ttl to now? does this mean on every request the reset time is pushed further out?

	remaining := rateLimitConfig.MaxRequests - int(count) // Why do we cast count to int here?
	if remaining < 0 {
		remaining = 0 // Are we just setting remaining to 0 if it's negative so we don't return a negative number in the header?
	}

	allowed := count <= int64(rateLimitConfig.MaxRequests)

	return allowed, remaining, resetTime, nil
}

func OpenAIRateLimiter() gin.HandlerFunc {
	return RateLimiter(DefaultOpenAIRateLimit)
}
