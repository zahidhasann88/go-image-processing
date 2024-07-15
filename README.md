Apologies for the confusion. Let's add Docker instructions specifically for your Golang Image Processing Library project:

Here's how you can modify your `README.md` file to include Docker instructions:

```markdown
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

## Docker Installation and Deployment

### Prerequisites
Make sure you have Docker installed on your machine. You can download it from [Docker's official website](https://www.docker.com/products/docker-desktop).

### Build Docker Image
To build the Docker image for the server, navigate to the root directory of your project where the Dockerfile is located and run:
```bash
docker build -t image-processing-server .
```

### Running the Docker Container
Once the image is built, you can run the server in a Docker container using:
```bash
docker run -p 8080:8080 -d image-processing-server
```
This command maps port 8080 of the Docker container to port 8080 on your host machine. Adjust the ports as needed if your server listens on a different port.

### Docker Compose
Alternatively, you can use Docker Compose for managing your application stack. Create a `docker-compose.yml` file in your project directory with the following content:
```yaml
version: '3'
services:
  image-processing:
    image: image-processing-server
    ports:
      - "8080:8080"
    environment:
      DATABASE_URL: "postgres://user:password@localhost/dbname?sslmode=disable"
      JWT_SECRET: "my_secret_key"
    depends_on:
      - postgres
  postgres:
    image: postgres
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: dbname
    ports:
      - "5432:5432"
```
This Docker Compose configuration sets up two services: `image-processing` for your Golang application and `postgres` for your PostgreSQL database. Adjust the environment variables and ports according to your setup.

## Using Makefile

## Endpoints
- Prerequisites
To use the Makefile on Windows, you need a tool that supports the make command. You can use one of the following methods:

1. Install Make for Windows:
- Download and install Make for Windows from GnuWin.
- Add the path to the make executable to your system's PATH environment variable.

2. Use Git Bash:
- Install Git for Windows from git-scm.com.
- Open Git Bash.

3.  Use Windows Subsystem for Linux (WSL):
Enable the Windows Subsystem for Linux and install a Linux distribution from the Microsoft Store.
Open your WSL terminal.

## Running Makefile Commands
1. Open your terminal (Git Bash, Command Prompt, or WSL).
2. Navigate to your project directory
```bash
cd path\to\your\project
```
3. Run the make command
```bash
make run
```

## Makefile Commands
- make run: Run the server.
- make test: Run the tests.
- make build: Build the application.
- make clean: Clean the build artifacts.

## Running the Server
If you prefer to run the server outside of Docker, you can still do so using:
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
```

This README.md file now includes Docker installation and deployment instructions specific to your Golang Image Processing Library project, alongside existing instructions for running the server, endpoints, configuration, and testing with Postman. Adjust paths and settings as necessary for your specific environment.