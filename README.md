# URL Shortener

> [🇧🇷 Versão em Português](README.pt-br.md)

A modern and efficient URL shortening service built with Go, using Gin Framework and PostgreSQL.

## 🚀 Features

- ✅ URL shortening with unique hash
- ✅ Automatic redirection
- ✅ Click counting
- ✅ RESTful API
- ✅ Swagger API documentation
- ✅ Structured logging with Zap
- ✅ Containerized with Docker
- ✅ PostgreSQL database
- ✅ Health check endpoint

## 🛠️ Tech Stack

- **Go 1.24.1** - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM for Go
- **PostgreSQL** - Database
- **Zap** - Structured logger
- **Swagger** - API documentation
- **Docker & Docker Compose** - Containerization

## 📁 Project Structure

```
url-shortener/
├── cmd/
│   └── shortener-api/
│       └── main.go              # Application entry point
├── docs/                        # Swagger documentation
│   ├── docs.go                  # Generated swagger docs
│   ├── swagger.json             # OpenAPI JSON specification
│   └── swagger.yaml             # OpenAPI YAML specification
├── internal/
│   ├── app/
│   │   ├── model/
│   │   │   └── Url.go           # URL data model
│   │   └── service/
│   │       └── shortener.go     # Business logic
│   ├── server/
│   │   ├── handlers/
│   │   │   └── url.go           # HTTP handlers
│   │   └── middleware/
│   │       └── logger.go        # Logging middleware
│   └── storage/
│       └── postgres/
│           └── repository.go    # Data repository
├── docker-compose.yml           # Container orchestration
├── dockerfile                   # Application Docker image
├── go.mod                       # Go dependencies
└── go.sum                       # Dependencies checksums
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24.1 or higher
- Docker and Docker Compose
- PostgreSQL (if running locally)

### Running with Docker Compose (Recommended)

1. Clone the repository:
```bash
git clone https://github.com/YuriGarciaRibeiro/url-shortener.git
cd url-shortener
```

2. Run with Docker Compose:
```bash
docker-compose up --build
```

The application will be available at `http://localhost:8080`

### 📚 API Documentation

The API documentation is available via Swagger UI at:
- **Swagger UI**: `http://localhost:8080/swagger/index.html`
- **OpenAPI JSON**: `http://localhost:8080/swagger/doc.json`
- **OpenAPI YAML**: Available in `docs/swagger.yaml`

### Running Locally

1. Clone the repository:
```bash
git clone https://github.com/YuriGarciaRibeiro/url-shortener.git
cd url-shortener
```

2. Set environment variables:
```bash
export DB_DSN="postgres://postgres:postgres@localhost:5432/url_shortener?sslmode=disable"
```

3. Install dependencies:
```bash
go mod download
```

4. Run the application:
```bash
go run cmd/shortener-api/main.go
```

## 📋 API Endpoints

> 💡 **Tip**: For complete API documentation with interactive examples, visit the Swagger UI at `http://localhost:8080/swagger/index.html` when the application is running.

### Health Check
```http
GET /health
```

**Response:**
```json
{
  "status": "ok"
}
```

### Shorten URL
```http
POST /shorten
Content-Type: application/json

{
  "url": "https://example.com/a-very-long-url"
}
```

**Response:**
```json
{
  "hash": "abc123",
  "short_url": "http://localhost:8080/abc123"
}
```

### Redirect
```http
GET /{hash}
```

Automatically redirects to the original URL and increments the click counter.

## 🗄️ Data Model

### URLs Table

| Field       | Type      | Description                  |
|-------------|-----------|------------------------------|
| ID          | uint      | Primary key                  |
| Hash        | string    | Unique hash for short URL    |
| OriginalURL | string    | Original URL                 |
| Clicks      | int       | Click counter                |
| CreatedAt   | timestamp | Creation date                |
| UpdatedAt   | timestamp | Last update date             |
| DeletedAt   | timestamp | Deletion date (soft delete)  |

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default Value |
|----------|-------------|---------------|
| `DB_DSN` | PostgreSQL connection string | `postgres://postgres:postgres@postgres:5432/url_shortener?sslmode=disable` |
| `PORT`   | Server port | `8080` |

## 🧪 Testing the API

### Using cURL

**Shorten a URL:**
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com/YuriGarciaRibeiro/url-shortener"}'
```

**Test redirection:**
```bash
curl -I http://localhost:8080/{returned_hash}
```

### Using HTTPie

**Shorten a URL:**
```bash
http POST localhost:8080/shorten url=https://github.com/YuriGarciaRibeiro/url-shortener
```

## 🏗️ Architecture

The project follows a clean architecture with clear separation of concerns:

- **cmd/**: Application entry point
- **internal/app/**: Business logic and models
- **internal/server/**: HTTP layer (handlers and middleware)
- **internal/storage/**: Persistence layer

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## 👨‍💻 Author

**Yuri Garcia Ribeiro**
- GitHub: [@YuriGarciaRibeiro](https://github.com/YuriGarciaRibeiro)

---

⭐ If this project helped you, consider giving it a star!
 
