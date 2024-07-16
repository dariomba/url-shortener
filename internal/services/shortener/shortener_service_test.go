package shortener_test

import (
	"testing"

	"github.com/dariomba/url-shortener/internal/services/shortener"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShortLink(t *testing.T) {
	tests := []struct {
		name         string
		originalURL  string
		expectedLink string
	}{
		{
			name:         "TestShortLinkWithValidURL",
			originalURL:  "https://github.com/dariomba/url-shortener/blob/master/cmd/main.go",
			expectedLink: "CkpsxkQq",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortLink, _ := shortener.GenerateShortLink(tt.originalURL)
			assert.Equal(t, tt.expectedLink, shortLink)
		})
	}
}
