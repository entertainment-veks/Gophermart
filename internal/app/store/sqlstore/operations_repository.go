package sqlstore

import (
	"gophermart/internal/app/model"
	"gophermart/internal/app/store"
)

type OperationsRepository struct {
	store *Store
}

func (r *OperationsRepository) Create(o *model.Operation) error {
	err := r.store.database.QueryRow(
		"INSERT INTO operations (owner, order_num, amount, processed_at) VALUES ($1, $2, $3, NOW()) RETURNING id",
		o.Owner,
		o.Order,
		o.Amount,
	).Scan(&o.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *OperationsRepository) GetBalance(login string) (int, error) {
	var balance *int
	err := r.store.database.QueryRow(
		"SELECT sum(amount) from operations where owner = $1",
		login,
	).Scan(&balance)

	if err != nil {
		return 0, err
	}

	return *balance, nil
}

func (r *OperationsRepository) GetWithdrawCount(login string) (int, error) {
	var count *int
	err := r.store.database.QueryRow(
		"SELECT count(amount) from operations where owner = $1 AND amount < 0",
		login,
	).Scan(&count)

	if err != nil {
		return 0, err
	}

	return *count, nil
}

func (r *OperationsRepository) GetAll(login string) ([]*model.Operation, error) {
	operations := []*model.Operation{}

	rows, err := r.store.database.Query(
		"SELECT order_num, amount, processed_at FROM operations WHERE owner = $1",
		login,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := &model.Operation{}
		err := rows.Scan(
			&o.Order,
			&o.Amount,
			&o.Processed_at,
		)
		if err != nil {
			return nil, err
		}
		operations = append(operations, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(operations) == 0 {
		return nil, store.ErrOperationsNotFound
	}

	return operations, nil
}
