package configs

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

var (
	RedisClient *redis.Client
	redisOnce   sync.Once
)

// Initialize Redis client once
//func init() {
//	redisOnce.Do(func() {
//		loadRedis()
//	})
//}

func loadRedis() {
	addr := fmt.Sprintf("%s:%s",
		getEnv("CACHE_HOST", "localhost"),
		getEnv("CACHE_PORT", "6379"),
	)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: getEnv("CACHE_PASS", ""),
		DB:       0,
	})

	_, err := RedisClient.Ping(RedisCtx()).Result()
	if err != nil {
		log.Fatalf("❌ Failed to connect to Redis: %v", err)
	}

	log.Println("✅ Redis connected:", addr)
}

func RedisCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

func ReloadRedis() {
	loadRedis()
	log.Println("✅ Redis config reloaded")
}
