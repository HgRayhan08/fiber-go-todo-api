package repository

import (
	"context"
	"database/sql"
	"errors"
	"todo-list/domain"
	"todo-list/dto"

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

// Show implements [domain.TaskRepository].
func (t *TodoDatabase) Show(ctx context.Context, idTask dto.IdTaskRequest) (result domain.Task, err error) {
	dataset := t.db.From("task").Where(goqu.C("id").Eq(idTask.Id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Create implements [domain.TaskRepository].
func (t *TodoDatabase) Create(ctx context.Context, todo domain.Task) error {
	executor := t.db.Insert("task").Rows(todo).Executor()
	_, err := executor.ExecContext(ctx)
	return err

}

// Delete implements [domain.TaskRepository].
func (t *TodoDatabase) Delete(ctx context.Context, idTodo string) error {
	dataset := t.db.Delete("task").Where(goqu.C("id").In(idTodo)).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// FindAll implements [domain.TaskRepository].
func (t *TodoDatabase) FindAll(ctx context.Context, id string) (result []domain.Task, err error) {

	if t.db == nil {
		return nil, errors.New("database is nil: goqu instance not initialized")
	}
	dataaset := t.db.From("task").Where(goqu.C("user_id").Eq(id))
	// dataaset := c.db.From("customers")
	err = dataaset.ScanStructsContext(ctx, &result)
	return

}

// FindById implements [domain.TaskRepository].
func (t *TodoDatabase) FindById(ctx context.Context, idTodo string) (result domain.Task, err error) {
	dataset := t.db.From("task").Where(goqu.C("id").Eq(idTodo))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Update implements [domain.TaskRepository].
func (t *TodoDatabase) Update(ctx context.Context, todo domain.Task) error {
	dataset := t.db.Update("task").Where(goqu.C("id").Eq(todo.Id)).Set(todo).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}
