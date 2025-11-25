package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

/*
*
Package-level Redis client variable
Similar to how we store the database connection, we store a Redis client
that can be shared across the application
*/
var RedisClient *redis.Client

func InitRedis() {
	// Get Redis URL from environment, with a sensible default for local dev
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
		log.Println("⚠️  REDIS_URL not set, using default:", redisAddr)
	}

	// Get optional password from environment
	redisPassword := os.Getenv("REDIS_PASSWORD")

	// Create Redis client
	// redis.NewClient returns a pointer to redis.Client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,     // Redis server address (host:port)
		Password: redisPassword, // Password (empty string if no password)
		DB:       0,             // Default DB (Redis has 16 databases by default, numbered 0-15)

		// Connection pool settings (similar to PostgreSQL pool)
		PoolSize:     10,              // Maximum number of socket connections
		MinIdleConns: 5,               // Minimum number of idle connections
		MaxIdleConns: 10,              // Maximum number of idle connections
		PoolTimeout:  4 * time.Second, // Time to wait for connection from pool

		// Timeouts
		DialTimeout:  5 * time.Second, // Timeout for establishing new connections
		ReadTimeout:  3 * time.Second, // Timeout for socket reads
		WriteTimeout: 3 * time.Second, // Timeout for socket writes
	})

	// Test the connection using a context with timeout
	// context.Background() creates an empty context
	// We add a 5-second timeout to avoid hanging forever
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Always clean up context resources

	// Ping Redis to verify connection
	// Ping() returns (string, error) - we only care about the error
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("✅ Redis connected successfully")
}

// GetRedis returns the Redis client instance
// Similar to GetDB(), this provides safe access to the Redis client
// with a nil check to prevent crashes
func GetRedis() *redis.Client {
	if RedisClient == nil {
		log.Fatal("Redis not initialized. Call InitRedis() first.")
	}
	return RedisClient
}

// CloseRedis closes the Redis connection
// Should be called when the application shuts down
func CloseRedis() error {
	if RedisClient != nil {
		return RedisClient.Close()
	}
	return nil
}
