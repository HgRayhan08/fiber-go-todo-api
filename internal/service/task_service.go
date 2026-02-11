package service

import (
	"context"
	"database/sql"
	"time"
	"todo-list/domain"
	"todo-list/dto"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type TodoService struct {
	todoRepository domain.TaskRepository
}

func NewTodoService(todoRepository domain.TaskRepository) domain.TaskService {
	return &TodoService{
		todoRepository: todoRepository,
	}
}

// Create implements [domain.TodoService].
func (t *TodoService) Create(ctx context.Context, f fiber.Ctx, request dto.CreateTaskRequest) error {
	userID := f.Locals("user_id").(string)
	todo := domain.Task{
		Id:          uuid.New().String(),
		UserID:      userID,
		Title:       request.Title,
		Description: request.Description,
		Status:      "Progress",
		CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}
	return t.todoRepository.Create(ctx, todo)
}

// Index implements [domain.TodoService].
func (t *TodoService) Index(ctx context.Context, f fiber.Ctx) ([]dto.TaskData, error) {
	userID := f.Locals("user_id").(string)
	todo, err := t.todoRepository.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	var formattedTodo []dto.TaskData

	for _, v := range todo {
		var updatedAt *time.Time = nil
		if v.UpdatedAt.Valid {
			t := v.UpdatedAt.Time
			updatedAt = &t
		}
		formattedTodo = append(formattedTodo, dto.TaskData{
			Id:          v.Id,
			UserID:      userID,
			Title:       v.Title,
			Description: v.Description,
			Status:      v.Status,
			CreatedAt:   v.CreatedAt.Time,
			UpdatedAt:   updatedAt,
		})

	}
	return formattedTodo, nil
}

// Delete implements [domain.TaskService].
func (t *TodoService) Delete(ctx context.Context, f fiber.Ctx, id string) error {
	return t.todoRepository.Delete(ctx, id)
}
