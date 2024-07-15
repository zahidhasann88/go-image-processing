# Golang Image Processing Library

## Overview
This library provides a set of image processing functions with a RESTful API interface.

## Features
- Resize
- Crop
- Rotate
- Blur
- Grayscale
- Sharpen
- Authentication with JWT
- Rate limiting

## Installation
```bash
go mod tidy
```
## Running the Server
```bash
go run cmd/server/main.go
```

### Endpoints
POST /register: Register a new user
POST /login: Login and get a token
POST /upload: Upload and process a single image
POST /batch-upload: Upload and process multiple images
POST /async-upload: Upload and process an image asynchronously

## Configuration
Create a .env file with the following content:
```bash
DATABASE_URL=postgres://user:password@localhost/dbname?sslmode=disable
JWT_SECRET=my_secret_key
```