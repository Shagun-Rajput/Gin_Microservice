#Gin_Microservice

This project is a starter template for building RESTful microservices using [Gin](https://github.com/gin-gonic/gin) in Go. It demonstrates a clean architecture with separate layers for configuration, routing, handlers, services, repositories, and models, and uses a PostgreSQL database.

## Features

- RESTful API with Gin
- PostgreSQL connection with connection pooling
- Environment-based configuration via `.env`
- Layered architecture (handler, service, repository, model)
- Example `/users` endpoint

## Project Structure

```
Gin_Microservice/
├── cmd/                # Application entry point
│   └── main.go
├── internal/
│   ├── config/         # Configuration and DB connection
│   ├── handler/        # HTTP handlers
│   ├── model/          # Data models (entities)
│   ├── repository/     # Database access logic
│   └── router/         # Route setup
├── .env                # Environment variables
├── go.mod
└── go.sum
```

## Getting Started

### 1. Clone the repository

```sh
git clone <your-repo-url>
cd Gin_Microservice
```

### 2. Set up environment variables

Edit the `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_SCHEMA=your_db_schema
```

### 3. Install dependencies

```sh
go mod tidy
```

### 4. Run the application

From the project root, run:

```sh
go run ./cmd
```

The server will start on port `8181` by default (or set `PORT` in your `.env`).

### 5. Example API Usage

- **GET /users**  
  Returns all users from the database.

  ```sh
  curl http://localhost:8181/users
  ```

## Notes

- The database connection uses `sslmode=disable` for development. Adjust as needed for production.
- Make sure your PostgreSQL instance is running and accessible with the credentials in `.env`.
- The default schema is set via the `search_path` parameter in the connection string.

## License