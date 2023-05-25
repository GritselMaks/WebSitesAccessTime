package interfaces

import (
	"context"
	"siteavliable/internal/models"
)

//go:generate mockgen -source=interfaces.go -destination=./../mocks_repo/mocks.go -package=mocks

type (

	// RedisRepoStats interface
	IRedisRepoStats interface {
		Save(context.Context, []models.CounterStats) error
		Get(context.Context, []string) ([]models.CounterStats, error)
	}

	// RedisRepoClients interface
	IRedisRepoClients interface {
		GetWithMax(ctx context.Context) (string, int64, error)
		GetWithMin(ctx context.Context) (string, int64, error)
		GetByURL(ctx context.Context, siteName string) (int64, error)
	}

	// RedisRepoClients interface
	IRedisRepoUpdate interface {
		SetByURL(ctx context.Context, siteName string, accessTime int64) error
	}
)
