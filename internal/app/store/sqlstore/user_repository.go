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

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		return store.ErrUserAlreadyExist
	}
	return nil
}

func (r *UserRepository) GetByLogin(login string) (*model.User, error) {
	u := &model.User{}
	err := r.store.database.QueryRow(
		"SELECT id, login, password FROM users WHERE login = $1",
		login,
	).Scan(
		&u.ID,
		&u.Login,
		&u.Password,
	)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, store.ErrUserNotFound
	}

	return u, nil
}
