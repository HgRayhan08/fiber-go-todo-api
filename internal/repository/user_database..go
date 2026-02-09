package repository

import (
	"context"
	"database/sql"
	"todo-list/domain"

	"github.com/doug-martin/goqu/v9"
)

type userDatabase struct {
	db *goqu.Database
}

func NewUserDatabase(con *sql.DB) domain.UserRepository {
	return &userDatabase{
		db: goqu.New("postgres", con),
	}
}

// Save implements [domain.UserRepository].
func (u *userDatabase) Save(ctx context.Context, user domain.User) error {
	dataset := u.db.Insert("users").Rows(user).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// FindByEmail implements [domain.UserRepository].
func (u *userDatabase) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("email").Eq(email))

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}
