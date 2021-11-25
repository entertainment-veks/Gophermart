package user

import (
	"encoding/json"
	"errors"
	"gophermart/internal/app/handler"
	"gophermart/internal/app/model"
	"gophermart/internal/app/service/user"
	"gophermart/internal/app/store"
	"net/http"
)

func RegisterHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u *model.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			handler.Error(w, http.StatusBadRequest, err)
			return
		}

		if err := user.Register(s, u); err != nil {
			switch {
			case errors.Is(err, handler.ErrBadRequest):
				handler.Error(w, http.StatusBadRequest, errors.Unwrap(err))
			case errors.Is(err, handler.ErrInternalServer):
				handler.Error(w, http.StatusInternalServerError, errors.Unwrap(err))
			case errors.Is(err, handler.ErrConflict):
				handler.Error(w, http.StatusConflict, errors.Unwrap(err))
			}
			return
		}

		authUser(w, u.Login)
		handler.Respond(w, http.StatusOK, "success")
	}
}
