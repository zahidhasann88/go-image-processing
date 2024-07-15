package processing

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"golang.org/x/image/bmp"
)

type Image struct {
	img image.Image
}

// LoadImage loads an image from a file
func LoadImage(filename string) (*Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	return &Image{img: img}, nil
}

// Resize resizes the image to the given width and height
func (img *Image) Resize(width, height int) {
	img.img = imaging.Resize(img.img, width, height, imaging.Lanczos)
}

// Crop crops the image to the given rectangle
func (img *Image) Crop(rect image.Rectangle) {
	img.img = imaging.Crop(img.img, rect)
}

// Rotate90 rotates the image by 90 degrees
func (img *Image) Rotate90() {
	img.img = imaging.Rotate90(img.img)
}

// Blur applies a blur effect to the image
func (img *Image) Blur(sigma float64) {
	img.img = imaging.Blur(img.img, sigma)
}

// Grayscale converts the image to grayscale
func (img *Image) Grayscale() {
	img.img = imaging.Grayscale(img.img)
}

// Sharpen applies a sharpen effect to the image
func (img *Image) Sharpen() {
	img.img = imaging.Sharpen(img.img, 0.5)
}

// SaveImage saves the image to a file
func (img *Image) SaveImage(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	switch ext := filepath.Ext(filename); ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img.img, nil)
	case ".png":
		err = png.Encode(file, img.img)
	case ".bmp":
		err = bmp.Encode(file, img.img)
	default:
		return fmt.Errorf("unsupported file format: %v", ext)
	}

	if err != nil {
		return fmt.Errorf("failed to encode image: %v", err)
	}

	return nil
}
