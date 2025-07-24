# Golang Application

This directory contains the source code and Dockerfile for the Go-based sample application.

## Structure
- `main.go`: Main application entry point.
- `main_test.go`: Unit tests for the application.
- `go.mod`: Go module definition.
- `Dockerfile`: Containerization instructions for building and running the Go app.

## Dockerfile
The `Dockerfile` is a multi-stage build that compiles the Go application and produces a minimal image for deployment. It uses the official Golang image for building and then copies the binary into a scratch or distroless image for production.

### Build and Run with Docker
```sh
# Build the Docker image
docker build -t devops-sample-apps-golang .

# Run the container
docker run --rm -p 8080:80 devops-sample-apps-golang
```

## Testing
Run tests locally with:
```sh
go test -v ./...
```

Or inside Docker:
```sh
docker run --rm devops-sample-apps-golang go test -v ./...
``` 
