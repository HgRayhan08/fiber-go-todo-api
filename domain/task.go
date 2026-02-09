package domain

import (
	"context"
	"database/sql"
	"todo-list/dto"
)

type Task struct {
	Id          string       `db:"id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	Status      string       `db:"status_id"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type TaskRepository interface {
	FindAll(ctx context.Context) ([]Task, error)
	FindById(ctx context.Context, idTask string) (Task, error)
	Create(ctx context.Context, task Task) error
	Update(ctx context.Context, task Task) error
	Delete(ctx context.Context, idTask string) error
}

type TaskService interface {
	Index(ctx context.Context) ([]dto.TaskData, error)
	Create(ctx context.Context, request dto.CreateTaskRequest) error
}
