package storage_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dariomba/url-shortener/src/internal/mocks"
	"github.com/dariomba/url-shortener/src/internal/services/storage"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

type mocksStorage struct {
	storageClient *mocks.MockStorageClient
}

func TestSaveURL(t *testing.T) {
	ctx := context.Background()

	shortURL := "http://example.com/jhdsjkfh3"
	ogURL := "http://original.url.domain.too.long.url/directory/other/files/example/file"

	tests := []struct {
		name        string
		shortURL    string
		originalURL string
		expectError bool
		mocks       func(m mocksStorage)
	}{
		{
			name:        "WhenSetFails_ThenReturnsError",
			shortURL:    shortURL,
			originalURL: ogURL,
			expectError: true,
			mocks: func(m mocksStorage) {
				statusCmd := redis.NewStatusCmd(ctx)
				statusCmd.SetErr(errors.New("weird error"))
				m.storageClient.EXPECT().Set(ctx, shortURL, ogURL, storage.CacheDuration).Return(statusCmd)
			},
		},
		{
			name:        "WhenEverythingOK_ThenReturnsNil",
			shortURL:    shortURL,
			originalURL: ogURL,
			expectError: false,
			mocks: func(m mocksStorage) {
				statusCmd := redis.NewStatusCmd(ctx)
				statusCmd.SetVal("OK")
				m.storageClient.EXPECT().Set(ctx, shortURL, ogURL, storage.CacheDuration).Return(statusCmd)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mocksStorage{
				storageClient: mocks.NewMockStorageClient(ctrl),
			}

			tt.mocks(m)

			service := storage.NewStorageService(m.storageClient)
			ctx := context.Background()

			err := service.SaveURL(ctx, tt.shortURL, tt.originalURL)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetURL(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		shortURL    string
		expectedURL string
		expectError bool
		mocks       func(m mocksStorage)
	}{
		{
			name:        "WhenGetAnExistingLink_ThenReturnFullURL",
			shortURL:    "short123",
			expectedURL: "http://original.url",
			expectError: false,
			mocks: func(m mocksStorage) {
				stringCmd := redis.NewStringCmd(ctx)
				stringCmd.SetVal("http://original.url")
				m.storageClient.EXPECT().Get(ctx, "short123").Return(stringCmd)
			},
		},
		{
			name:        "WhenGetANonExistingLink_ThenReturnsAnError",
			shortURL:    "nonexistent",
			expectedURL: "",
			expectError: true,
			mocks: func(m mocksStorage) {
				stringCmd := redis.NewStringCmd(ctx)
				stringCmd.SetErr(redis.Nil)
				m.storageClient.EXPECT().Get(ctx, "nonexistent").Return(stringCmd)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mocksStorage{
				storageClient: mocks.NewMockStorageClient(ctrl),
			}

			tt.mocks(m)

			service := storage.NewStorageService(m.storageClient)
			ctx := context.Background()

			retrievedURL, err := service.GetURL(ctx, tt.shortURL)
			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedURL, retrievedURL)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedURL, retrievedURL)
			}
		})
	}
}
