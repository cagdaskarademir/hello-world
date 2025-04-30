# Hello World API

A simple Go API with health checks and hello endpoint, built with Gin framework.

## Features

- Health check endpoints (liveness and readiness)
- Hello endpoint with JSON response
- Swagger documentation
- Docker support
- GitHub Actions CI/CD pipeline

## Prerequisites

- Go 1.23 or later
- Docker and Docker Compose
- Make (optional)

## Getting Started

### Local Development

1. Clone the repository:

```bash
git clone https://github.com/cagdaskarademir/hello-world.git
cd hello-world
```

2. Install dependencies:

```bash
go mod download
```

3. Run the application:

```bash
go run cmd/api/main.go
```

### Docker

1. Build and run with Docker Compose:

```bash
docker-compose up --build
```

## API Endpoints

- `GET /api/v1/health/liveness` - Liveness check
- `GET /api/v1/health/readiness` - Readiness check
- `GET /api/v1/say/:name` - Returns a greeting message

## Swagger Documentation

Access the Swagger documentation at:

```
http://localhost:8080/swagger/index.html
```

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   ├── health.go
│   │   └── hello.go
│   └── middleware/
│       └── middleware.go
├── docs/
├── pkg/
│   ├── health/
│   └── response/
├── .github/
│   └── workflows/
│       └── ci-cd.yml
├── Dockerfile
├── docker-compose.yml
└── README.md
```
