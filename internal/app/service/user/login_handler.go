package user

import (
	"encoding/json"
	"gophermart/internal/app/model"
	"gophermart/internal/app/service"
	"gophermart/internal/app/store"
	"net/http"
)

func LoginHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user *model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			service.Error(w, http.StatusBadRequest, err)
			return
		}

		if err := validate(user); err != nil {
			service.Error(w, http.StatusBadRequest, err)
			return
		}

		findedUser, err := s.User().GetByLogin(user.Login)
		if err != nil && err != store.ErrUserNotFound {
			service.Error(w, http.StatusInternalServerError, err)
			return
		}
		if err == store.ErrUserNotFound {
			service.Error(w, http.StatusUnauthorized, err)
			return
		}

		if err := hashPassword(user); err != nil {
			service.Error(w, http.StatusInternalServerError, err)
			return
		}

		if equal(user, findedUser) {
			authUser(w, user.Login)
			service.Respond(w, http.StatusOK, "success")
			return
		}
		service.Respond(w, http.StatusUnauthorized, "incorrect password")
	}
}
