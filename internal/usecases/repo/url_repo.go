package repo

import (
	"context"
	"errors"
	redisclient "siteavliable/pkg/client/redis"

	"github.com/redis/go-redis/v9"
)

var ErrValueNotFound = errors.New("value not found")

type urlRedisRepo struct {
	setName string
	r       *redisclient.RedisClient
}

// NewUrlsRepo returns a new urlRedisRepo instance
func NewUrlsRepo(c *redisclient.RedisClient, setname string) *urlRedisRepo {
	return &urlRedisRepo{setName: setname,
		r: c}
}

// GetWithMax returns an url, access time and error with maximum access time value
func (u *urlRedisRepo) GetWithMax(ctx context.Context) (string, int64, error) {
	res, err := u.r.Client.ZRevRangeWithScores(ctx, u.setName, 0, 0).Result()
	if err != nil {
		return "", 0, err
	}
	if len(res) > 0 {
		key := res[0].Member.(string)
		value := res[0].Score
		return key, int64(value), nil
	}
	return "", 0, ErrValueNotFound
}

// GetWithMin returns an url, access time and error with minimum access time value
func (u *urlRedisRepo) GetWithMin(ctx context.Context) (string, int64, error) {
	res, err := u.r.Client.ZRangeWithScores(ctx, u.setName, 0, 0).Result()
	if err != nil {
		return "", 0, err
	}
	if len(res) > 0 {
		key := res[0].Member.(string)
		value := res[0].Score
		return key, int64(value), nil
	}
	return "", 0, ErrValueNotFound
}

// GetByURL returns an access time and error by url
func (u *urlRedisRepo) GetByURL(ctx context.Context, url string) (int64, error) {
	res, err := u.r.Client.ZScore(ctx, u.setName, url).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, ErrValueNotFound
		}
		return 0, err
	}

	return int64(res), nil
}

// SetByURL add url and access time to storage
func (u *urlRedisRepo) SetByURL(ctx context.Context, url string, accessTime int64) error {
	err := u.r.Client.ZAdd(ctx, u.setName, redis.Z{
		Score:  float64(accessTime),
		Member: url,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}
