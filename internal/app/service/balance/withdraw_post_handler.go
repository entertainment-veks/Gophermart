package balance

import (
	"encoding/json"
	"gophermart/internal/app/handler"
	"gophermart/internal/app/handler/user"
	"gophermart/internal/app/model"
	"gophermart/internal/app/service/orders"
	"gophermart/internal/app/store"
	"net/http"
)

func WithdrawPostHandler(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var operation *model.Operation
		if err := json.NewDecoder(r.Body).Decode(&operation); err != nil {
			handler.Error(w, http.StatusBadRequest, err)
			return
		}

		if !orders.IsValid(operation.Order) {
			handler.Respond(w, http.StatusUnprocessableEntity, "invalid order number")
			return
		}

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

		if operation.Amount < 0 {
			handler.Respond(w, http.StatusBadRequest, "amount must be positive")
		}

		if balance < operation.Amount {
			handler.Respond(w, http.StatusPaymentRequired, "not enough funds")
			return
		}

		operation.Owner = login
		operation.Amount = operation.Amount * -1
		if err := store.Operations().Create(operation); err != nil {
			handler.Error(w, http.StatusInternalServerError, err)
		}
	}
}
