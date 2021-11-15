package sqlstore

import (
	"database/sql"
	"gophermart/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	database             *sql.DB
	userRepository       *UserRepository
	ordersRepository     *OrdersRepository
	operationsRepository *OperationsRepository
}

func New(database *sql.DB) *Store {
	return &Store{
		database: database,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}

func (s *Store) Orders() store.OrdersRepository {
	if s.ordersRepository == nil {
		s.ordersRepository = &OrdersRepository{
			store: s,
		}
	}

	return s.ordersRepository
}

func (s *Store) Operations() store.OperationsRepository {
	if s.operationsRepository == nil {
		s.operationsRepository = &OperationsRepository{
			store: s,
		}
	}

	return s.operationsRepository
}
