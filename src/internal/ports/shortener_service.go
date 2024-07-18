package ports

//go:generate mockgen -source=./shortener_service.go -destination=../mocks/shortener_service_mock.go -package=mocks
type ShortenerService interface {
	GenerateShortLink(originalURL string) (string, error)
}
