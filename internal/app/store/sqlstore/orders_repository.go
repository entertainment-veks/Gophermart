package sqlstore

import (
	"database/sql"
	"gophermart/internal/app/model"
	"gophermart/internal/app/store"
)

type OrdersRepository struct {
	store *Store
}

func (r *OrdersRepository) Create(o *model.Order) error {
	err := r.store.database.QueryRow(
		"INSERT INTO orders (number, status, owner) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id",
		o.Number,
		o.Status,
		o.Owner,
	).Scan(&o.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrdersRepository) GetOwnerByNumber(number int) (string, error) {
	var owner *string
	err := r.store.database.QueryRow(
		"SELECT owner FROM orders WHERE number = $1",
		number,
	).Scan(
		&owner,
	)

	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if err == sql.ErrNoRows {
		return "", store.ErrOrderNotExist
	}

	return *owner, nil
}
