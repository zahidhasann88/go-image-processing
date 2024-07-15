package main

import (
	"image-processing-library/api/handlers"
	"image-processing-library/pkg/auth"
	"image-processing-library/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize the database connection
	auth.InitDB("postgres://postgres:admin123@localhost/imgproc?sslmode=disable")

	// Authentication routes
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)

	// Image processing routes
	router.POST("/upload", auth.AuthMiddleware(), handlers.HandleUpload)
	router.POST("/batch-upload", auth.AuthMiddleware(), handlers.HandleBatchUpload)
	router.POST("/async-upload", auth.AuthMiddleware(), handlers.HandleAsyncUpload)

	// Rate limiting
	router.Use(middleware.RateLimiter(10))

	// Start the server
	router.Run(":8080")
}
