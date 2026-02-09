package domain

import (
	"context"
	"database/sql"
)

type Category struct {
	Id        string       `db:"id"`
	Name      string       `db:"name"`
	CreatedAt sql.NullTime `db:"created_at"`
}

type CategoryRepository interface {
	FindAllUser(ctx context.Context, idUser string) ([]Category, error)
	Save(ctx context.Context, category Category) error
	Delete(ctx context.Context, idCategory string) error
}

type CategoryService interface {
	IndexUser(ctx context.Context, idUser string) ([]Category, error)
	Create(ctx context.Context, name string) error
	Delete(ctx context.Context, idCategory string) error
}
