package store

import "gophermart/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	GetByLogin(string) (*model.User, error)
}
