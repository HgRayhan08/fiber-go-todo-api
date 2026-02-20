package repository

import (
	"context"
	"database/sql"
	"todo-list/domain"

	"github.com/doug-martin/goqu/v9"
)

type CategoryDatabase struct {
	db *goqu.Database
}

func NewCategoryDatabase(con *sql.DB) domain.CategoryRepository {
	return &CategoryDatabase{
		db: goqu.New("postgres", con),
	}
}

// FindById implements [domain.CategoryRepository].
func (c *CategoryDatabase) FindById(ctx context.Context, idCategory string) (result domain.Category, err error) {
	dataset := c.db.From("category").Where(goqu.C("id").Eq(idCategory))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Delete implements [domain.CategoryRepository].
func (c *CategoryDatabase) Delete(ctx context.Context, idCategory string) error {
	dataset := c.db.Delete("category").Where(goqu.C("id").In(idCategory)).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// FindAllUser implements [domain.CategoryRepository].
func (c *CategoryDatabase) FindAllUser(ctx context.Context, idUser string) (result []domain.Category, err error) {
	database := c.db.From("category").Where(goqu.C("user_id").Eq(idUser))
	err = database.ScanStructsContext(ctx, &result)
	return
}

// Save implements [domain.CategoryRepository].
func (c *CategoryDatabase) Save(ctx context.Context, category domain.Category) error {
	database := c.db.Insert("category").Rows(category).Executor()
	_, err := database.ExecContext(ctx)
	return err
}
