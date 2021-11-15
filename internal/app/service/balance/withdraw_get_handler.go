package balance

import (
	"gophermart/internal/app/service"
	"gophermart/internal/app/service/user"
	"gophermart/internal/app/store"
	"net/http"
)

func WithdrawGetHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login, err := user.GetLogin(r)
		if err != nil {
			service.Error(w, http.StatusInternalServerError, err)
			return
		}

		oprations, err := s.Operations().GetAll(login)
		if err != nil && err != store.ErrOperationsNotFound {
			service.Error(w, http.StatusInternalServerError, err)
			return
		}
		if err == store.ErrOperationsNotFound {
			service.Respond(w, http.StatusNoContent, "no operations was found")
			return
		}

		service.RespondJSON(w, http.StatusOK, oprations)
	}
}
