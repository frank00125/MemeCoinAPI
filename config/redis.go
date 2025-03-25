package config

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewRedisClient() (*redis.Client, error) {
	redisHost := viper.GetString("REDIS_HOST")
	redisPassword := viper.GetString("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
