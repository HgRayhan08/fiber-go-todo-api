package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todo-list/domain"
	"todo-list/dto"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type TodoService struct {
	todoRepository     domain.TaskRepository
	categoryRepository domain.CategoryRepository
}

func NewTodoService(todoRepository domain.TaskRepository, categoryRepository domain.CategoryRepository) domain.TaskService {
	return &TodoService{
		todoRepository:     todoRepository,
		categoryRepository: categoryRepository,
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
	fmt.Printf("ini data ", todo)
	var formattedTodo []dto.TaskData

	for _, v := range todo {

		categoryData, err := t.categoryRepository.FindById(ctx, v.CategoryID)
		if err != nil {
			return nil, err
		}

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
			Category:    categoryData.Name,
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
