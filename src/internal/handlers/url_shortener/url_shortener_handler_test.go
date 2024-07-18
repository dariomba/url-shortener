package urlshortener_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	urlshortener "github.com/dariomba/url-shortener/src/internal/handlers/url_shortener"
	"github.com/dariomba/url-shortener/src/internal/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mocksShortenerHandler struct {
	storageService   *mocks.MockStorageService
	shortenerService *mocks.MockShortenerService
}

func TestCreateLink(t *testing.T) {
	type want struct {
		statusCode int
		body       map[string]string
	}

	tests := []struct {
		name        string
		want        want
		requestBody map[string]string
		mocks       func(m mocksShortenerHandler)
	}{
		{
			name:        "WhenCreatesALinkWithEmptyURL_ThenReturnsBadRequest",
			requestBody: map[string]string{},
			want:        want{statusCode: http.StatusBadRequest, body: map[string]string{"error": "url parameter is required"}},
			mocks:       func(m mocksShortenerHandler) {},
		},
		{
			name:        "WhenGenerateShortLinkFails_ThenReturnsInternalServerError",
			requestBody: map[string]string{"url": "http://example.com"},
			want:        want{statusCode: http.StatusInternalServerError, body: map[string]string{"error": "an error has ocurred creating the link"}},
			mocks: func(m mocksShortenerHandler) {
				m.shortenerService.EXPECT().GenerateShortLink("http://example.com").Return("", errors.New("new error"))
			},
		},
		{
			name:        "WhenSaveURLFails_ThenReturnsInternalServerError",
			requestBody: map[string]string{"url": "http://example.com"},
			want:        want{statusCode: http.StatusInternalServerError, body: map[string]string{"error": "an error has ocurred creating the link"}},
			mocks: func(m mocksShortenerHandler) {
				m.shortenerService.EXPECT().GenerateShortLink("http://example.com").Return("shortLink", nil)
				m.storageService.EXPECT().SaveURL(gomock.Any(), "shortLink", "http://example.com").Return(errors.New("new error"))
			},
		},
		{
			name:        "WhenEverythingOK_ThenReturnsFullShortURL",
			requestBody: map[string]string{"url": "http://example.com"},
			want: want{statusCode: http.StatusOK, body: map[string]string{"message": "short url created successfully!",
				"url": "http://localhost/shortLink"}},
			mocks: func(m mocksShortenerHandler) {
				m.shortenerService.EXPECT().GenerateShortLink("http://example.com").Return("shortLink", nil)
				m.storageService.EXPECT().SaveURL(gomock.Any(), "shortLink", "http://example.com").Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mocksShortenerHandler{
				storageService:   mocks.NewMockStorageService(ctrl),
				shortenerService: mocks.NewMockShortenerService(ctrl),
			}

			tt.mocks(m)

			gin.SetMode(gin.TestMode)
			router := gin.Default()
			group := router.Group("/")

			os.Setenv("HOST", "http://localhost/")
			urlshortener.NewURLShortenerHandler(group, m.storageService, m.shortenerService)

			w := httptest.NewRecorder()
			body, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/createLink", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(w, req)

			var response map[string]string
			json.Unmarshal(w.Body.Bytes(), &response)

			assert.Equal(t, tt.want.statusCode, w.Code)
			assert.Equal(t, tt.want.body, response)
		})
	}
}

func TestRedirectToURL(t *testing.T) {
	type want struct {
		statusCode int
		body       map[string]string
		URL        string
	}

	tests := []struct {
		name  string
		want  want
		link  string
		mocks func(m mocksShortenerHandler)
	}{
		{
			name: "WhenGetURLFails_ThenReturnsNotFound",
			link: "noExists",
			want: want{statusCode: http.StatusNotFound, body: map[string]string{"error": "URL not found"}},
			mocks: func(m mocksShortenerHandler) {
				m.storageService.EXPECT().GetURL(gomock.Any(), "noExists").Return("", errors.New("not found"))
			},
		},
		{
			name: "WhenEverythingOK_ThenRedirectsToURL",
			link: "someLink",
			want: want{statusCode: http.StatusFound, URL: "http://example.com"},
			mocks: func(m mocksShortenerHandler) {
				m.storageService.EXPECT().GetURL(gomock.Any(), "someLink").Return("http://example.com", nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mocksShortenerHandler{
				storageService:   mocks.NewMockStorageService(ctrl),
				shortenerService: mocks.NewMockShortenerService(ctrl),
			}

			tt.mocks(m)

			gin.SetMode(gin.TestMode)
			router := gin.Default()
			group := router.Group("/")

			urlshortener.NewURLShortenerHandler(group, m.storageService, m.shortenerService)

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", "/"+tt.link, nil)

			router.ServeHTTP(w, req)

			var response map[string]string
			json.Unmarshal(w.Body.Bytes(), &response)

			assert.Equal(t, tt.want.statusCode, w.Code)
			if tt.want.body != nil {
				var response map[string]string
				json.Unmarshal(w.Body.Bytes(), &response)
			}
			if tt.want.URL != "" {
				assert.Equal(t, tt.want.URL, w.Header().Get("Location"))
			}
		})
	}
}
