package domain

import (
	"context"
	"database/sql"
	"todo-list/dto"

	"github.com/gofiber/fiber/v3"
)

type Category struct {
	Id        string       `db:"id"`
	Name      string       `db:"name"`
	UserId    string       `db:"user_id"`
	CreatedAt sql.NullTime `db:"created_at"`
}

type CategoryRepository interface {
	FindAllUser(ctx context.Context, idUser string) ([]Category, error)
	Save(ctx context.Context, category Category) error
	Delete(ctx context.Context, idCategory string) error
}

type CategoryService interface {
	IndexUser(ctx context.Context, f fiber.Ctx) ([]dto.CategoryData, error)
	Create(ctx context.Context, f fiber.Ctx, name dto.CreateCategoryRequest) error
	Delete(ctx context.Context, f fiber.Ctx, idCategory string) error
}
