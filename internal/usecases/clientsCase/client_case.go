package clientscase

import (
	"context"
	"log"
	"siteavliable/internal/models"
	"siteavliable/internal/usecases/interfaces"
)

// ClientUseCase-...
type ClientUseCase struct {
	logger *log.Logger
	repo interfaces.IRedisRepoClients
}

// New returns a new ClientUseCase instance
func New(r interfaces.IRedisRepoClients, l *log.Logger) *ClientUseCase {
	return &ClientUseCase{
		logger: l,
		repo: r,
	}
}

// GetWithMinResponeTime returns an url with the minimum access time
func (cc *ClientUseCase) GetWithMinResponeTime(ctx context.Context) (models.AccessTime, error) {
	url, accessTime, err := cc.repo.GetWithMin(ctx)
	if err != nil {
		return models.AccessTime{}, err
	}
	return models.AccessTime{URL: url, AccessTime: accessTime}, nil
}

// GetWithMaxResponeTime returns an url with the maximum access time
func (cc *ClientUseCase) GetWithMaxResponeTime(ctx context.Context) (models.AccessTime, error) {
	url, accessTime, err := cc.repo.GetWithMax(ctx)
	if err != nil {
		return models.AccessTime{}, err
	}
	return models.AccessTime{URL: url, AccessTime: accessTime}, nil
}

// GetByURL returns an access time by url
func (cc *ClientUseCase) GetByURL(ctx context.Context, url string) (models.AccessTime, error) {
	accessTime, err := cc.repo.GetByURL(ctx, url)
	if err != nil {
		return models.AccessTime{}, err
	}
	return models.AccessTime{URL: url, AccessTime: accessTime}, nil
}
