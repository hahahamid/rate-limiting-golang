package main

import (
	"rate-limiting-demo-go/handlers"
	"rate-limiting-demo-go/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func main() {
	r := gin.Default()

	limiter := middleware.NewIPRateLimiter(rate.Every(1*time.Minute), 100)

	api := r.Group("/api")
	api.Use(middleware.RateLimitMiddleware(limiter))
	{
		api.GET("/message", handlers.GetMessage)
	}

	r.Run(":8080")
}
