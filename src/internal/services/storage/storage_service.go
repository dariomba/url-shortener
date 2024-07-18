package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/dariomba/url-shortener/src/internal/ports"
)

const CacheDuration = 8 * time.Hour

type StorageService struct {
	client ports.StorageClient
}

func NewStorageService(client ports.StorageClient) *StorageService {
	return &StorageService{
		client: client,
	}
}

func (s StorageService) SaveURL(ctx context.Context, shortURL string, originalURL string) error {
	err := s.client.Set(ctx, shortURL, originalURL, CacheDuration).Err()
	if err != nil {
		return fmt.Errorf("an error has ocurred saving the url --> %w", err)
	}
	return nil
}

func (s StorageService) GetURL(ctx context.Context, shortURL string) (string, error) {
	url, err := s.client.Get(ctx, shortURL).Result()
	if err != nil {
		return "", fmt.Errorf("an error has ocurred retrieving the url --> %w", err)
	}
	return url, nil
}
