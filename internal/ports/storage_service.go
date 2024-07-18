package ports

import "context"

type StorageService interface {
	SaveURL(ctx context.Context, shortURL string, originalURL string) error
	GetURL(ctx context.Context, shortURL string) (string, error)
}
