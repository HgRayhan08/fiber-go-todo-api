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

	// middleware JWT Protected
	jwtMidd := middleware.JWTProtected(*config)

	// Middleware Rate Limiting
	app.Use(middleware.RateLimited())

	// repository
	taskDatabase := repository.NewTodoDatabase(dbConnection)
	authDatabase := repository.NewUserDatabase(dbConnection)
	categoryDatabase := repository.NewCategoryDatabase(dbConnection)

	// Service
	TaskService := service.NewTodoService(taskDatabase, categoryDatabase)
	AuthService := service.NewAuthService(config, authDatabase)
	categoryService := service.NewCategoryService(categoryDatabase)

	// API
	api.NewAuthApi(app, AuthService)
	api.NewTaskApi(app, TaskService, jwtMidd)
	api.NewCategoryApi(app, categoryService, jwtMidd)

	// api.NewTodoDatabase(todoDatabase)
	_ = app.Listen(config.Server.Host + ":" + config.Server.Port)
}
