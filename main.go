package main

import (
	"todo-list/internal/api"
	"todo-list/internal/config"
	"todo-list/internal/connection"
	"todo-list/internal/middleware"
	"todo-list/internal/repository"
	"todo-list/internal/service"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config := config.Get()

	// Connection Database
	dbConnection := connection.GetDatabaseConnection(config.Database)

	// Initialize Fiber app
	app := fiber.New(
		fiber.Config{
			ErrorHandler: middleware.ErrorHandler(),
		},
	)
	jwtMidd := middleware.JWTProtected(*config)
	// Middleware Rate Limiting
	app.Use(middleware.RateLimited())

	// repository
	taskDatabase := repository.NewTodoDatabase(dbConnection)
	authDatabase := repository.NewUserDatabase(dbConnection)

	// Service
	TaskService := service.NewTodoService(taskDatabase)
	AuthService := service.NewAuthService(config, authDatabase)

	// API
	api.NewTaskApi(app, TaskService, jwtMidd)
	api.NewAuthApi(app, AuthService)

	// api.NewTodoDatabase(todoDatabase)
	_ = app.Listen(config.Server.Host + ":" + config.Server.Port)
}
