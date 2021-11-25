package balance

import (
	"gophermart/internal/app/handler"
	"gophermart/internal/app/handler/user"
	"gophermart/internal/app/store"
	"net/http"
)

func BalanceHandler(store store.Store) http.HandlerFunc {
	type responseJSON struct {
		Current   int `json:"current"`
		Withdrawn int `json:"withdrawn"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		login, err := user.GetLogin(r)
		if err != nil {
			handler.Error(w, http.StatusInternalServerError, err)
			return
		}

		balance, err := store.Operations().GetBalance(login)
		if err != nil {
			handler.Error(w, http.StatusInternalServerError, err)
			return
		}

		withdrawn, err := store.Operations().GetWithdrawCount(login)
		if err != nil {
			handler.Error(w, http.StatusInternalServerError, err)
			return
		}

		response := &responseJSON{
			Current:   balance,
			Withdrawn: withdrawn,
		}

		handler.RespondJSON(w, http.StatusOK, response)
	}
}
