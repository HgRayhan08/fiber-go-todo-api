package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"todo-list/domain"

	"github.com/doug-martin/goqu/v9"
)

type TodoDatabase struct {
	db *goqu.Database
}

func NewTodoDatabase(con *sql.DB) domain.TaskRepository {
	return &TodoDatabase{
		db: goqu.New("postgres", con),
	}
}

// Create implements [domain.TaskRepository].
func (t *TodoDatabase) Create(ctx context.Context, todo domain.Task) error {
	executor := t.db.Insert("task").Rows(todo).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Delete implements [domain.TaskRepository].
func (t *TodoDatabase) Delete(ctx context.Context, idTodo string) error {
	panic("unimplemented")
}

// FindAll implements [domain.TaskRepository].
func (t *TodoDatabase) FindAll(ctx context.Context, id string) (result []domain.Task, err error) {

	if t.db == nil {
		fmt.Println("database is nil: goqu instance not initialized")
		return nil, errors.New("database is nil: goqu instance not initialized")
	}
	dataaset := t.db.From("task").Where(goqu.C("user_id").Eq(id))
	// dataaset := c.db.From("customers")
	err = dataaset.ScanStructsContext(ctx, &result)
	return

}

// FindById implements [domain.TaskRepository].
func (t *TodoDatabase) FindById(ctx context.Context, idTodo string) (result domain.Task, err error) {
	panic("unimplemented")
}

// Update implements [domain.TaskRepository].
func (t *TodoDatabase) Update(ctx context.Context, todo domain.Task) error {
	panic("unimplemented")
}
