package orders

import (
	"fmt"
	"gophermart/internal/app/handler"
	"gophermart/internal/app/model"
	"gophermart/internal/app/store"
	"log"
	"sort"
)

func GetAllByLogin(s store.Store, login string) ([]*model.Order, error) {
	orders, err := s.Orders().GetAllByUser(login)
	if err != nil && err != store.ErrOrdersNotFound {
		log.Println(err)
		return nil, fmt.Errorf("%w: %v", handler.ErrInternalServer, err)
	}
	if err == store.ErrOrdersNotFound {
		log.Println(err)
		return nil, fmt.Errorf("%w: %v", handler.ErrNoContent, err)
	}

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].Uploaded_at.Before(orders[j].Uploaded_at)
	})

	return orders, err
}
