package urlshortener

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/dariomba/url-shortener/internal/services/shortener"
	"github.com/dariomba/url-shortener/internal/services/storage"
	"github.com/gin-gonic/gin"
)

type URLShortenerHandler struct {
	storageService storage.StorageService
}

type CreateLinkRequest struct {
	URL string `json:"url" binding:"required"`
}

func NewURLShortenerHandler(
	router *gin.RouterGroup,
	storageService storage.StorageService,
) {
	urlShortenerHandler := URLShortenerHandler{
		storageService: storageService,
	}

	router.POST("/createLink", urlShortenerHandler.CreateLink)
	router.GET("/:link", urlShortenerHandler.RedirectToURL)
}

func (u *URLShortenerHandler) CreateLink(c *gin.Context) {
	var createLinkReq CreateLinkRequest
	if err := c.ShouldBindJSON(&createLinkReq); err != nil {
		log.Error(fmt.Errorf("binding the JSON --> %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "url parameter is required"})
		return
	}

	shortLink, err := shortener.GenerateShortLink(createLinkReq.URL)
	if err != nil {
		log.Error(fmt.Errorf("generating the link --> %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "an error has ocurred creating the link"})
		return
	}

	err = u.storageService.SaveURL(c, shortLink, createLinkReq.URL)
	if err != nil {
		log.Error(fmt.Errorf("saving the url --> %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "an error has ocurred creating the link"})
		return
	}

	host := "http://localhost:8080/"

	c.JSON(http.StatusOK, gin.H{
		"message": "short url created successfully!",
		"url":     host + shortLink,
	})
}

func (u *URLShortenerHandler) RedirectToURL(c *gin.Context) {
	link := c.Param("link")

	originalURL, err := u.storageService.GetURL(c, link)
	if err != nil {
		log.Error(fmt.Errorf("retrieving the original url --> %w", err))
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "URL not found"})

		return
	}

	c.Redirect(302, originalURL)
}
