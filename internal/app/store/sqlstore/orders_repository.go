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
		"INSERT INTO orders (number, status, owner, uploaded_at) VALUES ($1, $2, $3, NOW()) ON CONFLICT DO NOTHING RETURNING id",
		o.Number,
		o.Status,
		o.Owner,
	).Scan(&o.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrdersRepository) GetOwnerByNumber(number string) (string, error) {
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

func (r *OrdersRepository) GetAllByUser(login string) ([]*model.Order, error) {
	orders := []*model.Order{}

	rows, err := r.store.database.Query(
		"SELECT number, status, accrual, uploaded_at FROM orders WHERE owner = $1",
		login,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := &model.Order{}
		err := rows.Scan(
			&o.Number,
			&o.Status,
			&o.Accrual,
			&o.Uploaded_at,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, store.ErrOrdersNotFound
	}

	return orders, nil
}
