package store

import "gophermart/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	GetByLogin(string) (*model.User, error)
}

type OrdersRepository interface {
	Create(*model.Order) error
	GetOwnerByNumber(string) (string, error)
	GetAllByUser(string) ([]*model.Order, error)
	UpdateStatus(*model.Order) error
}

type OperationsRepository interface {
	Create(*model.Operation) error
	GetBalance(string) (int, error)
	GetWithdrawCount(string) (int, error)
	GetAll(string) ([]*model.Operation, error)
}
