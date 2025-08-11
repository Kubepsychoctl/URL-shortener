# URL Shortener

A high-performance URL shortening service built with Go, featuring authentication, click tracking, and real-time statistics.

## Features

- **URL Shortening**: Generate short, unique hashes for long URLs
- **Authentication**: JWT-based user authentication with registration/login
- **Click Tracking**: Monitor link performance with detailed analytics
- **Real-time Statistics**: Group click data by day or month
- **Event-driven Architecture**: Asynchronous click processing
- **RESTful API**: Clean, intuitive HTTP endpoints
- **PostgreSQL**: Robust database with GORM ORM
- **Docker Support**: Easy deployment with Docker Compose

## Architecture

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Client    │───▶│   Router    │───▶│  Handlers   │
└─────────────┘    └─────────────┘    └─────────────┘
                           │                   │
                    ┌─────────────┐    ┌─────────────┐
                    │ Middleware  │    │  Services   │
                    └─────────────┘    └─────────────┘
                           │                   │
                    ┌─────────────┐    ┌─────────────┐
                    │ Event Bus   │    │ Repository  │
                    └─────────────┘    └─────────────┘
                           │                   │
                           └─────────┐         │
                                     ▼         ▼
                              ┌─────────────┐ ┌─────────────┐
                              │ Background  │ │ PostgreSQL  │
                              │  Workers    │ │   Database  │
                              └─────────────┘ └─────────────┘
```

## Tech Stack

- **Language**: Go 1.24.1
- **Framework**: Standard `net/http` with custom middleware
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT tokens
- **Validation**: go-playground/validator
- **Testing**: Standard Go testing package
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Go 1.24.1 or higher
- PostgreSQL 16.4
- Docker & Docker Compose (optional)

## Quick Start

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/Kubepsychoctl/URL-shortener/
   cd URL-shortener
   ```

2. **Create environment file**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start services**
   ```bash
   docker-compose up -d
   ```

4. **Run migrations**
   ```bash
   go run migrations/auto.go
   ```

5. **Start the application**
   ```bash
   go run cmd/main.go
   ```

### Manual Setup

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Set up PostgreSQL database**
   ```bash
   # Create database and user
   createdb url_shortener
   ```

3. **Configure environment variables**
   ```bash
   export DSN="host=localhost user=postgres password=password dbname=url_shortener port=5432 sslmode=disable"
   export SECRET_KEY="your-secret-key-here"
   export POSTGRES_USER="postgres"
   export POSTGRES_PASSWORD="password"
   ```

4. **Run migrations**
   ```bash
   go run migrations/auto.go
   ```

5. **Start the server**
   ```bash
   go run cmd/main.go
   ```

The server will start on `http://localhost:8080`

## Configuration

Create a `.env` file in the root directory:

```env
# Database
DSN=host=localhost user=postgres password=password dbname=url_shortener port=5432 sslmode=disable

# Authentication
SECRET_KEY=your-super-secret-jwt-key-here

# PostgreSQL (for Docker)
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
```

## API Reference

### Authentication

#### Register User
```http
POST /auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword",
  "name": "John Doe"
}
```

#### Login
```http
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword"
}
```

### Links

#### Create Short Link
```http
POST /link
Content-Type: application/json

{
  "url": "https://example.com/very-long-url-that-needs-shortening"
}
```

#### Update Link
```http
PATCH /link/{id}
Authorization: Bearer <jwt-token>
Content-Type: application/json

{
  "url": "https://new-url.com",
  "hash": "newhash"
}
```

#### Delete Link
```http
DELETE /link/{id}
```

#### Get All Links
```http
GET /link?limit=10&offset=0
Authorization: Bearer <jwt-token>
```

#### Redirect to Original URL
```http
GET /{hash}
```

### Statistics

#### Get Click Statistics
```http
GET /stat?from=2024-01-01&to=2024-12-31&by=day
Authorization: Bearer <jwt-token>
```

**Query Parameters:**
- `from`: Start date (YYYY-MM-DD)
- `to`: End date (YYYY-MM-DD)
- `by`: Grouping (`day` or `month`)

**Response:**
```json
[
  {
    "period": "2024-01-01",
    "sum": 150
  },
  {
    "period": "2024-01-02",
    "sum": 89
  }
]
```

## Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run specific test
go test ./cmd/auth_test.go

# Run with verbose output
go test -v ./...
```

### E2E Tests

The project includes end-to-end tests that spin up the full application:

```bash
go test ./cmd/auth_test.go -v
```

## Database Schema

### Users
- `id`: Primary key
- `email`: Unique email address
- `password`: Hashed password
- `name`: User's full name
- `created_at`, `updated_at`, `deleted_at`: Timestamps

### Links
- `id`: Primary key
- `url`: Original long URL
- `hash`: Unique short hash
- `created_at`, `updated_at`, `deleted_at`: Timestamps

### Stats
- `id`: Primary key
- `link_id`: Foreign key to links table
- `clicks`: Number of clicks for the date
- `date`: Date of the clicks
- `created_at`, `updated_at`, `deleted_at`: Timestamps

## Event System

The application uses an event-driven architecture for click tracking:

1. **Link Visit**: When a user visits a short URL, a `LinkVisitedEvent` is published
2. **Background Processing**: A background worker consumes these events
3. **Statistics Update**: Click counts are updated in the database
4. **Real-time Data**: Statistics are immediately available via the API

## Deployment

### Production Considerations

1. **Environment Variables**: Use strong, unique secrets
2. **Database**: Consider connection pooling and read replicas
3. **HTTPS**: Use a reverse proxy (nginx/traefik) with SSL termination
4. **Monitoring**: Add health checks and metrics
5. **Logging**: Implement structured logging for production

### Docker Production

```bash
# Build production image
docker build -t url-shortener:latest .

# Run with production environment
docker run -d \
  -p 8080:8080 \
  -e DSN="your-production-dsn" \
  -e SECRET_KEY="your-production-secret" \
  url-shortener:latest
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Built with Go
