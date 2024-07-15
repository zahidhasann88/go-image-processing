package handlers

import (
	"fmt"
	"image"
	"image-processing-library/internal/processing"
	"mime/multipart"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type ProcessingRequest struct {
	Resize    string  `form:"resize"`
	Crop      string  `form:"crop"`
	Rotate    bool    `form:"rotate"`
	Blur      float64 `form:"blur"`
	Grayscale bool    `form:"grayscale"`
	Sharpen   bool    `form:"sharpen"`
}

var taskQueue chan Task

type Task struct {
	File *multipart.FileHeader
	Req  ProcessingRequest
}

func HandleUpload(c *gin.Context) {
	var req ProcessingRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Create a temporary file to store the uploaded image
	tempFile, err := os.CreateTemp("", "upload-*.jpg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary file"})
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Save the uploaded file to the temporary file
	if err := c.SaveUploadedFile(file, tempFile.Name()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
		return
	}

	// Load the image using our processing library
	img, err := processing.LoadImage(tempFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load image"})
		return
	}

	// Apply processing operations
	if req.Resize != "" {
		var width, height int
		if _, err := fmt.Sscanf(req.Resize, "%dx%d", &width, &height); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid resize format"})
			return
		}
		img.Resize(width, height)
	}

	if req.Crop != "" {
		var x, y, w, h int
		if _, err := fmt.Sscanf(req.Crop, "%d,%d,%d,%d", &x, &y, &w, &h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crop format"})
			return
		}
		img.Crop(image.Rect(x, y, x+w, y+h))
	}

	if req.Rotate {
		img.Rotate90()
	}

	if req.Blur > 0 {
		img.Blur(req.Blur)
	}

	if req.Grayscale {
		img.Grayscale()
	}

	if req.Sharpen {
		img.Sharpen()
	}

	// Save the processed image to a new file
	processedFile, err := os.CreateTemp("", "processed-*.jpg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary file for processed image"})
		return
	}
	defer os.Remove(processedFile.Name())
	defer processedFile.Close()

	if err := img.SaveImage(processedFile.Name()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save processed image"})
		return
	}

	// Return the processed image URL
	processedURL := fmt.Sprintf("http://localhost:8080/uploads/%s", processedFile.Name())
	c.JSON(http.StatusOK, gin.H{"message": "Image processed successfully", "url": processedURL})
}

func HandleBatchUpload(c *gin.Context) {
	var req ProcessingRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	files := form.File["images"]
	var wg sync.WaitGroup
	errors := make(chan error, len(files))
	var processedImagePaths []string

	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			defer wg.Done()

			tempFile, err := os.CreateTemp("", "upload-*.jpg")
			if err != nil {
				errors <- fmt.Errorf("failed to create temporary file: %w", err)
				return
			}
			defer os.Remove(tempFile.Name())
			defer tempFile.Close()

			if err := c.SaveUploadedFile(file, tempFile.Name()); err != nil {
				errors <- fmt.Errorf("failed to save uploaded file: %w", err)
				return
			}

			img, err := processing.LoadImage(tempFile.Name())
			if err != nil {
				errors <- fmt.Errorf("failed to load image: %w", err)
				return
			}

			// Apply processing operations
			if req.Resize != "" {
				var width, height int
				if _, err := fmt.Sscanf(req.Resize, "%dx%d", &width, &height); err != nil {
					errors <- fmt.Errorf("invalid resize format: %w", err)
					return
				}
				img.Resize(width, height)
			}

			if req.Crop != "" {
				var x, y, w, h int
				if _, err := fmt.Sscanf(req.Crop, "%d,%d,%d,%d", &x, &y, &w, &h); err != nil {
					errors <- fmt.Errorf("invalid crop format: %w", err)
					return
				}
				img.Crop(image.Rect(x, y, x+w, y+h))
			}

			if req.Rotate {
				img.Rotate90()
			}

			if req.Blur > 0 {
				img.Blur(req.Blur)
			}

			if req.Grayscale {
				img.Grayscale()
			}

			if req.Sharpen {
				img.Sharpen()
			}

			processedFile, err := os.CreateTemp("", "processed-*.jpg")
			if err != nil {
				errors <- fmt.Errorf("failed to create temporary file for processed image: %w", err)
				return
			}
			defer os.Remove(processedFile.Name())
			defer processedFile.Close()

			if err := img.SaveImage(processedFile.Name()); err != nil {
				errors <- fmt.Errorf("failed to save processed image: %w", err)
				return
			}
			// Save processedFile to a storage or return URLs
			processedImagePaths = append(processedImagePaths, processedFile.Name())
		}(file)
	}

	wg.Wait()
	close(errors)

	if len(errors) > 0 {
		var errorMessages []string
		for err := range errors {
			errorMessages = append(errorMessages, err.Error())
		}
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errorMessages})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":              "Batch processing completed successfully",
		"processed_image_uris": processedImagePaths,
	})
}

func HandleAsyncUpload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	var req ProcessingRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	taskQueue <- Task{File: file, Req: req}
	c.JSON(http.StatusAccepted, gin.H{"message": "Task accepted"})
}
