package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"go-rest-api/pkg/db"
)

const (
	selectAll = `SELECT * FROM users`
)

type User struct {
	ID          string    `db:"id" json:"id"`
	Name        string    `db:"user_name" validate:"required" json:"user_name"`
	Designation string    `db:"designation" validate:"required" json:"designation"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type UsersRepository interface {
	GetUsers(ctx context.Context) (users []User, err error)
}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetUsers(ctx context.Context) (users []User, err error) {
	dbInstance := db.GetDB()

	err = dbInstance.SelectContext(ctx, &users, selectAll)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "Failed to get all users")
	}

	return users, nil
}
