Todo List

A simple Golang REST API application for managing tasks with JWT authentication.

## **Tech Stack**
- [Golang](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [JWT](https://github.com/golang-jwt/jwt)

## **Folder Structure**
- Todo-List/
- ├── domain/               # File Model
- ├── dto/                  # Data Transfer Objects (DTO) untuk request & response
- ├── internal/
- │   ├── Api/              # API handlers & route definitions 
- |   ├── Config/           # File konfigurasi init database
- |   ├── Connection/       # File Database connection setup
- │   ├── middleware/       # JWT middleware, dll
- │   ├── repository/       # Logic Database accsess
- │   ├── service/          # Business logic
- │   └── utils/            # Helper functions and utilities
- ├── go.mod
- ├── go.sum
- └── main.go