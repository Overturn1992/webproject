package config

import (
	"backend/global"
	"github.com/go-redis/redis"
	"log"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Addr,
		DB:       AppConfig.Redis.Type,
		Password: AppConfig.Redis.Password,
	})
	if _, err := RedisClient.Ping().Result(); err != nil {
		log.Fatalf("Failed to connect to Redis,got error: %v", err)
	}
	global.RedisDB = RedisClient
}
