package ports

import "context"

//go:generate mockgen -source=./storage_service.go -destination=../mocks/storage_service_mock.go -package=mocks
type StorageService interface {
	SaveURL(ctx context.Context, shortURL string, originalURL string) error
	GetURL(ctx context.Context, shortURL string) (string, error)
}
