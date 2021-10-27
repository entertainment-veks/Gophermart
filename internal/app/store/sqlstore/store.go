package sqlstore

import (
	"database/sql"
	"gophermart/internal/app/store"
)

type Store struct {
	database       *sql.DB
	userRepository *UserRepository
}

func New(database *sql.DB) Store {
	return Store{
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
