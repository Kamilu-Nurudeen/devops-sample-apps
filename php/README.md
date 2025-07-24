# PHP Application

This directory contains the source code and Dockerfile for the PHP-based sample application.

## Structure
- `index.php`: Main PHP application file.
- `config.dev`, `config.prod`: Environment-specific configuration files.
- `Dockerfile`: Containerization instructions for building and running the PHP app.

## Dockerfile
The `Dockerfile` sets up a lightweight PHP environment using an official PHP base image. It copies the application code and configuration files into the container and exposes the necessary port for serving the application.

### Build and Run with Docker
```sh
# Build the Docker image
docker build -t devops-sample-apps-php .

# Run the container
docker run --rm -p 8080:80 devops-sample-apps-php
```

## Testing
You can check PHP syntax locally with:
```sh
php -l index.php
```

Or run the container and access the application via `http://localhost:8080`. 
