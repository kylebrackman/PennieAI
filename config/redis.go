package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
		log.Println("⚠️  REDIS_URL not set, using default:", redisAddr)
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

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

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	log.Println("✅ Redis connected successfully")
	return nil
}

func GetRedis() *redis.Client {
	if RedisClient == nil {
		log.Fatal("Redis not initialized. Call InitRedis() first.")
	}
	return RedisClient
}

// Should be called when the application shuts down
func CloseRedis() error {
	if RedisClient != nil {
		return RedisClient.Close()
	}
	return nil
}
