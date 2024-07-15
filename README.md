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

## Endpoints

### Authentication
- **POST /register**: Register a new user
  - **Request Body**:
    ```json
    {
      "username": "example",
      "password": "password123"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "User registered successfully"
    }
    ```

- **POST /login**: Login and get a token
  - **Request Body**:
    ```json
    {
      "username": "example",
      "password": "password123"
    }
    ```
  - **Response**:
    ```json
    {
      "token": "your_jwt_token"
    }
    ```

### Image Processing
- **POST /upload**: Upload and process a single image
  - **Authorization**: Bearer Token
  - **Request Body (form-data)**:
    - **Key**: `image` | **Value**: Select a file | **Type**: File
    - Additional keys (optional): `resize`, `crop`, `rotate`, `blur`, `grayscale`, `sharpen`
  - **Response**:
    ```json
    {
      "message": "Image processed successfully",
      "url": "http://localhost:8080/uploads/processed-image.jpg"
    }
    ```

- **POST /batch-upload**: Upload and process multiple images
  - **Authorization**: Bearer Token
  - **Request Body (form-data)**:
    - **Key**: `images` | **Value**: Select multiple files | **Type**: File
    - Additional keys (optional): `resize`, `crop`, `rotate`, `blur`, `grayscale`, `sharpen`
  - **Response**:
    ```json
    {
      "message": "Batch processing completed successfully"
    }
    ```

- **POST /async-upload**: Upload and process images asynchronously
  - **Authorization**: Bearer Token
  - **Request Body (form-data)**:
    - **Key**: `images` | **Value**: Select multiple files | **Type**: File
    - Additional keys (optional): `resize`, `crop`, `rotate`, `blur`, `grayscale`, `sharpen`
  - **Response**:
    ```json
    {
      "message": "Images are being processed"
    }
    ```

## Configuration
### Create a .env file with the following content:
```bash
DATABASE_URL=postgres://user:password@localhost/dbname?sslmode=disable
JWT_SECRET=my_secret_key
```

### Create a Postgres Database (Example Name: imgproc)
#### Run the SQL query
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

## Testing with Postman

### 1. Register a New User
- **Method**: POST
- **URL**: `http://localhost:8080/register`
- **Body**: 
  ```json
  {
    "username": "example",
    "password": "password123"
  }
  ```

### 2. Login and Get a Token
- **Method**: POST
- **URL**: `http://localhost:8080/login`
- **Body**: 
  ```json
  {
    "username": "example",
    "password": "password123"
  }
  ```
- **Response**:
  ```json
  {
    "token": "your_jwt_token"
  }
  ```

### 3. Upload and Process a Single Image
- **Method**: POST
- **URL**: `http://localhost:8080/upload`
- **Authorization**: Bearer Token
- **Body**: 
  - **Type**: form-data
  - **Key**: `image` | **Value**: Select a file | **Type**: File
  - Additional keys (optional): `resize`, `crop`, `rotate`, `blur`, `grayscale`, `sharpen`

### 4. Batch Upload and Process Multiple Images
- **Method**: POST
- **URL**: `http://localhost:8080/batch-upload`
- **Authorization**: Bearer Token
- **Body**: 
  - **Type**: form-data
  - **Key**: `images` | **Value**: Select multiple files | **Type**: File
  - Additional keys (optional): `resize`, `crop`, `rotate`, `blur`, `grayscale`, `sharpen`

### 5. Asynchronous Upload and Process Images
- **Method**: POST
- **URL**: `http://localhost:8080/async-upload`
- **Authorization**: Bearer Token
- **Body**: 
  - **Type**: form-data
  - **Key**: `images` | **Value**: Select multiple files | **Type**: File
  - Additional keys (optional): `resize`, `crop`, `rotate`, `blur`, `grayscale`, `sharpen`