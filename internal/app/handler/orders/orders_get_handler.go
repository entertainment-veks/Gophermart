package orders

import (
	"errors"
	"gophermart/internal/app/handler"
	"gophermart/internal/app/handler/user"
	"gophermart/internal/app/service/orders"
	"gophermart/internal/app/store"
	"net/http"
)

func OrdersGetHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userLogin, err := user.GetLogin(r)
		if err != nil {
			handler.Error(w, http.StatusInternalServerError, err)
		}

		allOrders, err := orders.GetAllByLogin(s, userLogin)
		if err != nil {
			switch {
			case errors.Is(err, handler.ErrNoContent):
				handler.Error(w, http.StatusNoContent, errors.Unwrap(err))
				break
			case errors.Is(err, handler.ErrInternalServer):
				handler.Error(w, http.StatusInternalServerError, errors.Unwrap(err))
				break
			}
		}

		handler.RespondJSON(w, http.StatusOK, allOrders)
	}
}
