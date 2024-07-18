package main

import (
	"fmt"
	"os"
	"strconv"

	urlshortener "github.com/dariomba/url-shortener/src/internal/handlers/url_shortener"
	"github.com/dariomba/url-shortener/src/internal/services/shortener"
	"github.com/dariomba/url-shortener/src/internal/services/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	v1 := router.Group("/")
	urlshortener.NewURLShortenerHandler(v1, *storage.NewStorageService(buildRedisClient()), &shortener.ShortenerService{})

	err := router.Run(":8080") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(fmt.Errorf("failed to start web server -> %w", err))
	}
}

func buildRedisClient() *redis.Client {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic("Redis DB must be an integer")
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PSWD"),
		DB:       redisDB,
	})
	return redisClient
}
