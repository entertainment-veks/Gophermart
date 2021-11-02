package orders

import (
	"gophermart/internal/app/service"
	"gophermart/internal/app/service/user"
	"gophermart/internal/app/store"
	"net/http"
)

func OrdersGetHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie(user.AuthCookieKey)
		if err != nil {
			service.Error(w, http.StatusInternalServerError, err) //here user cookie must exist
		}

		orders, err := s.Orders().GetAllByUser(authCookie.Value)
		if err != nil && err != store.ErrOrdersNotFound {
			service.Error(w, http.StatusInternalServerError, err)
			return
		}
		if err == store.ErrOrdersNotFound {
			service.Error(w, http.StatusNoContent, err)
			return
		}

		for _, currentOrder := range orders {
			if currentOrder.AccrualFromDB.Valid {
				currentOrder.Accrual = int(currentOrder.AccrualFromDB.Int16)
			}
		}

		service.RespondJSON(w, http.StatusOK, orders)
	}
}
