package balance

import (
	"gophermart/internal/app/store"
	"net/http"
)

func WithdrawPostHandler(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//
	}
}
