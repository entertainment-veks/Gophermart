package sqlstore

import "gophermart/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(*model.User) error {
	return nil
}
