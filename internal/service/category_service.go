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

type CategoryService struct {
	categoryRepository domain.CategoryRepository
}

func NewCategoryService(categoryRepository domain.CategoryRepository) domain.CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

// IndexById implements [domain.CategoryService].
func (c *CategoryService) IndexById(ctx context.Context, idCategory string) (dto.CategoryData, error) {
	category, err := c.categoryRepository.FindById(ctx, idCategory)
	if err != nil {
		return dto.CategoryData{}, err
	}
	return dto.CategoryData{
		Id:        category.Id,
		Name:      category.Name,
		UserID:    category.UserId,
		CreatedAt: category.CreatedAt.Time.String(),
	}, nil
}

// Create implements [domain.CategoryService].
func (c *CategoryService) Create(ctx context.Context, f fiber.Ctx, data dto.CreateCategoryRequest) error {
	userID := f.Locals("user_id").(string)
	category := domain.Category{
		Id:        uuid.New().String(),
		Name:      data.Name,
		UserId:    userID,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	return c.categoryRepository.Save(ctx, category)

}

// Delete implements [domain.CategoryService].
func (c *CategoryService) Delete(ctx context.Context, f fiber.Ctx, idCategory string) error {
	return c.categoryRepository.Delete(ctx, idCategory)
}

// IndexUser implements [domain.CategoryService].
func (c *CategoryService) IndexUser(ctx context.Context, f fiber.Ctx) ([]dto.CategoryData, error) {

	userId := f.Locals("user_id").(string)
	data, err := c.categoryRepository.FindAllUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	var categoryUser []dto.CategoryData

	for _, v := range data {
		if v.UserId != userId {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized access to category")
		}

		categoryUser = append(categoryUser, dto.CategoryData{
			Id:        v.Id,
			Name:      v.Name,
			UserID:    v.UserId,
			CreatedAt: v.CreatedAt.Time.String(),
		})
	}
	return categoryUser, nil
}
