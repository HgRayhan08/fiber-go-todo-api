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
	todoRepository     domain.TaskRepository
	categoryRepository domain.CategoryRepository
}

func NewTodoService(todoRepository domain.TaskRepository, categoryRepository domain.CategoryRepository) domain.TaskService {
	return &TodoService{
		todoRepository:     todoRepository,
		categoryRepository: categoryRepository,
	}
}

// Show implements [domain.TaskService].
func (t *TodoService) Show(ctx context.Context, idtask dto.IdTaskRequest) (dto.TaskData, error) {
	dataTask, err := t.todoRepository.Show(ctx, idtask)
	if err != nil {
		return dto.TaskData{}, err
	}

	var createdAt time.Time
	if dataTask.CreatedAt.Valid {
		t := dataTask.UpdatedAt.Time
		createdAt = t
	}

	var updatedAt *time.Time = nil
	if dataTask.UpdatedAt.Valid {
		t := dataTask.UpdatedAt.Time
		updatedAt = &t
	}
	return dto.TaskData{
		Id:          dataTask.Id,
		UserID:      dataTask.UserID,
		Title:       dataTask.Title,
		Description: dataTask.Description,
		Category:    dataTask.Category,
		CategoryId:  dataTask.CategoryID,
		Status:      dataTask.Status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

// Update implements [domain.TaskService].
func (t *TodoService) Update(ctx context.Context, f fiber.Ctx, request domain.Task) error {
	dataTask, err := t.todoRepository.FindById(ctx, request.Id)
	if err != nil {
		return err
	}

	CategoryData, err := t.categoryRepository.FindById(ctx, request.CategoryID)
	if err != nil {
		return err
	}

	dataTask.Title = request.Title
	dataTask.Description = request.Description
	dataTask.Status = request.Status
	dataTask.CategoryID = request.CategoryID
	dataTask.Category = CategoryData.Name
	dataTask.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return t.todoRepository.Update(ctx, dataTask)
}

// Create implements [domain.TodoService].
func (t *TodoService) Create(ctx context.Context, f fiber.Ctx, request dto.TaskRequest) error {
	userID := f.Locals("user_id").(string)

	categoryData, err := t.categoryRepository.FindById(ctx, request.CategoryID)

	if err != nil {
		return err
	}

	todo := domain.Task{
		Id:          uuid.New().String(),
		UserID:      userID,
		Title:       request.Title,
		CategoryID:  request.CategoryID,
		Category:    categoryData.Name,
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
			CategoryId:  v.CategoryID,
			Category:    v.Category,
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
