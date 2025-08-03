# Booking App - Backend

## 🚀 Run Locally

### ✅ Prerequisites

Make sure you have the following installed on your system:

- [Go](https://go.dev/doc/install) (v1.20 or later recommended)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Goose](https://github.com/pressly/goose) (for database migrations)
- [Postman](https://www.postman.com/downloads/) (for API testing)

To install Goose:

```bash
  go install github.com/pressly/goose/v3/cmd/goose@latest
```

Clone the project:

```bash
  git clone https://github.com/Kenasvarghese/Booking-App.git
```

Navigate to the project directory:

```bash
  cd Booking-App/Backend
```

Start PostgreSQL using Docker:

```bash
  docker compose up -d
```

Run database migration:

```bash
  goose up
```

Source environment variables

```bash
  source .env.sample
```

Start the server:

```bash
  go run main.go
```

## 🧪 API Testing with Postman

A ready-to-use Postman collection is included to test all API endpoints.

### 📂 Collection Location

```bash
  Backend/postman/booking-app.postman_collection.json
```

### 🔧 How to Use

1. Open [Postman](https://www.postman.com/)
2. Click **Import** → Upload the collection file from the path above
3. Set the `base_url` variable to:

```bash
  http://localhost:8080/api
```


     