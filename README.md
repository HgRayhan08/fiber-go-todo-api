Todo List

A simple Golang REST API application for managing tasks with JWT authentication.

## **Tech Stack**
- [Golang](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [JWT](https://github.com/golang-jwt/jwt)

## **Folder Structure**
Todo-list
├── domain/ # Domain models / Entity definitions
├── dto/ # Data Transfer Objects (DTO) for request & response
├── internal/
│ ├── api/ # API handlers & route definitions
│ ├── config/ # App configuration (e.g., database initialization)
│ ├── connection/ # Database connection setup
│ ├── middleware/ # JWT, logging, and other middleware
│ ├── repository/ # Database access / repository layer
│ ├── service/ # Business logic
│ └── utils/ # Helper functions and utilities
├── go.mod
├── go.sum
└── main.go