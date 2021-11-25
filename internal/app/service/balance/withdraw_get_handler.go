package balance

import (
	"gophermart/internal/app/handler"
	"gophermart/internal/app/handler/user"
	"gophermart/internal/app/store"
	"net/http"
)

func WithdrawGetHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login, err := user.GetLogin(r)
		if err != nil {
			handler.Error(w, http.StatusInternalServerError, err)
			return
		}

		oprations, err := s.Operations().GetAll(login)
		if err != nil && err != store.ErrOperationsNotFound {
			handler.Error(w, http.StatusInternalServerError, err)
			return
		}
		if err == store.ErrOperationsNotFound {
			handler.Respond(w, http.StatusNoContent, "no operations was found")
			return
		}

		handler.RespondJSON(w, http.StatusOK, oprations)
	}
}
