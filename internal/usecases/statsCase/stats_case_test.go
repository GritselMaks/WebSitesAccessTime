package statscase

import (
	"context"
	"errors"
	"log"
	"siteavliable/internal/metrics"
	"siteavliable/internal/models"
	mocks "siteavliable/internal/usecases/mocks_repo"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestStatsUseCase_SaveMetrics(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIRedisRepoStats(ctl)
	logger := log.Default()
	ctx := context.Background()
	m := []string{metrics.GetWithMaxTime, metrics.GetWithMinTime, metrics.GetByURL}
	uCase := New(repo, logger, m)
	metrics.Init()
	metrics.IncCounter(metrics.GetWithMaxTime)
	metrics.IncCounter(metrics.GetWithMinTime)
	metrics.IncCounter(metrics.GetByURL)
	stats := []models.CounterStats{
		{
			Handler: metrics.GetWithMaxTime,
			Counter: 1,
		},
		{
			Handler: metrics.GetWithMinTime,
			Counter: 1,
		},
		{
			Handler: metrics.GetByURL,
			Counter: 1,
		},
	}
	repo.EXPECT().Save(ctx, stats).Return(nil).Times(1)
	uCase.SaveMetrics(ctx)

	metrics.IncCounter(metrics.GetWithMaxTime)
	metrics.IncCounter(metrics.GetWithMinTime)
	metrics.IncCounter(metrics.GetByURL)

	errorRedis := errors.New("connection timeout errors")
	repo.EXPECT().Save(ctx, stats).Return(errorRedis).Times(1)
	uCase.SaveMetrics(ctx)
}
func TestStatsUseCase_GetMetrics(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIRedisRepoStats(ctl)
	logger := log.Default()
	ctx := context.Background()
	m := []string{metrics.GetWithMaxTime, metrics.GetWithMinTime, metrics.GetByURL}
	uCase := New(repo, logger, m)
	metrics.Init()
	metrics.IncCounter(metrics.GetWithMaxTime)
	metrics.IncCounter(metrics.GetWithMinTime)
	metrics.IncCounter(metrics.GetByURL)
	stats := []models.CounterStats{
		{
			Handler: metrics.GetWithMaxTime,
			Counter: 1,
		},
		{
			Handler: metrics.GetWithMinTime,
			Counter: 1,
		},
		{
			Handler: metrics.GetByURL,
			Counter: 1,
		},
	}
	repo.EXPECT().Get(ctx, uCase.metrics).Return(stats, nil).Times(1)
	s, err := uCase.GetMetrics(ctx)
	if err != nil {
		t.Errorf("should not return an error")
	}
	_ = s
	errorRedis := errors.New("connection timeout errors")
	repo.EXPECT().Get(ctx, uCase.metrics).Return(nil, errorRedis).Times(1)
	s, err = uCase.GetMetrics(ctx)
	if err == nil {
		t.Errorf("should not return an error")
	}
	if s != nil {
		t.Errorf("should return nil value")
	}
}
