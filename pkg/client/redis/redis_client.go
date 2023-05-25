package redisclient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Config - ...
type Config struct {
	Addr string
	DB   int
	Set  string
}

// RedisClient -...
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient return Red
func NewRedisClient(ctx context.Context, cfg Config) (*RedisClient, error) {
	c := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		DB:   cfg.DB,
	})
	_, err := c.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	client := RedisClient{
		Client: c,
	}
	return &client, err
}
