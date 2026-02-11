package domain

import (
	"context"
	"database/sql"
	"todo-list/dto"

	"github.com/gofiber/fiber/v3"
)

type Task struct {
	Id          string       `db:"id"`
	UserID      string       `db:"user_id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	CategoryID  string       `db:"category_id"`
	Category    string       `db:"category"`
	Status      string       `db:"status"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type TaskRepository interface {
	FindAll(ctx context.Context, idUser string) ([]Task, error)
	FindById(ctx context.Context, idTask string) (Task, error)
	Show(ctx context.Context, idTask dto.IdTaskRequest) (Task, error)
	Create(ctx context.Context, task Task) error
	Update(ctx context.Context, task Task) error
	Delete(ctx context.Context, idTask string) error
}

type TaskService interface {
	Index(ctx context.Context, f fiber.Ctx) ([]dto.TaskData, error)
	Show(ctx context.Context, idtask dto.IdTaskRequest) (dto.TaskData, error)
	Create(ctx context.Context, f fiber.Ctx, request dto.TaskRequest) error
	Update(ctx context.Context, f fiber.Ctx, request Task) error
	Delete(ctx context.Context, f fiber.Ctx, id string) error
}
