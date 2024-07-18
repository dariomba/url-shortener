package ports

type ShortenerService interface {
	GenerateShortLink(originalURL string) (string, error)
}
