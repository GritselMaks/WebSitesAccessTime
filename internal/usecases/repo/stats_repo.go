package repo

import (
	"context"
	"fmt"
	"siteavliable/internal/models"
	redisclient "siteavliable/pkg/client/redis"
)

type statRedisRepo struct {
	client *redisclient.RedisClient
}

// NewSatsRepo returns a statRedisRepo instance
func NewSatsRepo(c *redisclient.RedisClient) *statRedisRepo {
	return &statRedisRepo{client: c}
}

// Save saves a statistics in storage
func (s *statRedisRepo) Save(ctx context.Context, stats []models.CounterStats) error {
	for _, stat := range stats {
		if err := s.client.Client.IncrBy(ctx, stat.Handler, stat.Counter).Err(); err != nil {
			return fmt.Errorf("error redis save: %s", err.Error())
		}
	}
	return nil
}

func (s *statRedisRepo) Get(ctx context.Context, keys []string) ([]models.CounterStats, error) {
	stas := make([]models.CounterStats, len(keys))
	for i, k := range keys {
		res, err := s.client.Client.Get(ctx, k).Int64()
		if err != nil {
			return nil, err
		}
		stas[i] = models.CounterStats{
			Handler: k,
			Counter: res,
		}
	}
	return stas, nil
}
