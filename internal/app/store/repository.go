package store

import "gophermart/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	GetByLogin(string) (*model.User, error)
}

type OrdersRepository interface {
	Create(*model.Order) error
	GetOwnerByNumber(int) (string, error)
	GetAllByUser(string) ([]*model.Order, error)
}
