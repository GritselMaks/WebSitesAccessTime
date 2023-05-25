package redisclient

import (
	"context"
	"testing"
)

// TestClient return client with real connection and func for clearing db
func TestClient(t *testing.T, config Config) (*RedisClient, func()) {
	ctx := context.Background()
	t.Helper()
	client, err := NewRedisClient(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	teardown := func() {
		client.Client.FlushDB(ctx)
	}
	return client, teardown
}
