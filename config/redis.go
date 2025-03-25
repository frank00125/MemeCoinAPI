package config

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewRedisClient() (*redis.Client, error) {
	redisUrl := viper.GetString("REDIS_URL")
	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
