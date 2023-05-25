package statscase

import (
	"context"
	"log"
	"siteavliable/internal/metrics"
	"siteavliable/internal/models"
	"siteavliable/internal/usecases/interfaces"
)

// StatsUseCase -...
type StatsUseCase struct {
	metrics []string
	repo    interfaces.IRedisRepoStats
	logger  *log.Logger
}

// New returns a new StatsUseCase instance
func New(r interfaces.IRedisRepoStats, l *log.Logger, metrics []string) *StatsUseCase {
	return &StatsUseCase{
		metrics: metrics,
		repo:    r,
		logger:  l,
	}
}

// SaveMetrics get actualy metrics and save them in storage
func (u *StatsUseCase) SaveMetrics(ctx context.Context) {
	stats := make([]models.CounterStats, len(u.metrics))
	for i, v := range u.metrics {
		count := metrics.GetCounterAndClear(v)
		stats[i] = models.CounterStats{
			Handler: v,
			Counter: count,
		}
	}
	err := u.repo.Save(ctx, stats)
	if err != nil {
		u.logger.Printf("error save metrics; error: %s\n", err.Error())
	}
}

// GetMetrics get metrics from storage
func (u *StatsUseCase) GetMetrics(ctx context.Context) ([]models.CounterStats, error) {
	stats, err := u.repo.Get(ctx, u.metrics)
	if err != nil {
		u.logger.Printf("error get metrics; error: %s\n", err.Error())
		return nil, err
	}
	return stats, nil
}
