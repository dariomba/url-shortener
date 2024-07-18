package ports

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:generate mockgen -source=./storage_client.go -destination=../mocks/storage_client_mock.go -package=mocks
type StorageClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}
