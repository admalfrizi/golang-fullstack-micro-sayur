package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "redis-user:6379",
	})

	_, err := client.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	return client
}