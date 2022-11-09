package datastore

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
	"gitlab.com/aksaratech/barber-backend/infrastructure/env"
)

func LoadRedis() *redis.Client {
	r := env.ENV.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       r.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Error("redis connect ping failed, err:", err)
		return nil
	}

	fmt.Println("redis connect ping response:", pong)
	return client
}
