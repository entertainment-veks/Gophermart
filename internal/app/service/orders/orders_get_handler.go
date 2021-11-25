package orders

import (
	"gophermart/internal/app/handler"
	"gophermart/internal/app/handler/user"
	"gophermart/internal/app/store"
	"net/http"
	"sort"
)

func OrdersGetHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userLogin, err := user.GetLogin(r)
		if err != nil {
			handler.Error(w, http.StatusInternalServerError, err)
		}

		orders, err := s.Orders().GetAllByUser(userLogin)
		if err != nil && err != store.ErrOrdersNotFound {
			handler.Error(w, http.StatusInternalServerError, err)
			return
		}
		if err == store.ErrOrdersNotFound {
			handler.Error(w, http.StatusNoContent, err)
			return
		}

		sort.Slice(orders, func(i, j int) bool {
			return orders[i].Uploaded_at.Before(orders[j].Uploaded_at)
		})

		handler.RespondJSON(w, http.StatusOK, orders)
	}
}
