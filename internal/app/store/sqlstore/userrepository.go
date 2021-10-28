package sqlstore

import (
	"database/sql"
	"gophermart/internal/app/model"
	"gophermart/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	err := r.store.database.QueryRow(
		"INSERT INTO users (login, password) VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING id",
		u.Login,
		u.Password,
	).Scan(&u.ID)

	switch err {
	case sql.ErrNoRows:
		return store.ErrUserAlreadyExist
	case nil:
		return nil
	default:
		return err
	}
}
