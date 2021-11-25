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

func LoginHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u *model.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			handler.Error(w, http.StatusBadRequest, err)
			return
		}

		ok, err := user.Login(s, u)
		if err != nil {
			switch {
			case errors.Is(err, handler.ErrBadRequest):
				handler.Error(w, http.StatusBadRequest, errors.Unwrap(err))
				break
			case errors.Is(err, handler.ErrInternalServer):
				handler.Error(w, http.StatusInternalServerError, errors.Unwrap(err))
				break
			case errors.Is(err, handler.ErrUnauthorized):
				handler.Error(w, http.StatusUnauthorized, errors.Unwrap(err))
				break
			}
		}

		if ok {
			authUser(w, u.Login)
			handler.Respond(w, http.StatusOK, "success")
			return
		}

		handler.Respond(w, http.StatusUnauthorized, "incorrect password")
	}
}
