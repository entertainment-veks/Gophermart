package orders

import (
	"gophermart/internal/app/service"
	"gophermart/internal/app/service/user"
	"gophermart/internal/app/store"
	"net/http"
	"sort"
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

		sort.Slice(orders, func(i, j int) bool {
			return orders[i].Uploaded_at.Before(orders[j].Uploaded_at)
		})

		service.RespondJSON(w, http.StatusOK, orders)
	}
}
