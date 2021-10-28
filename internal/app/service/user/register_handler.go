package user

import (
	"encoding/json"
	"gophermart/internal/app/model"
	"gophermart/internal/app/service"
	"gophermart/internal/app/store"
	"net/http"
)

func RegisterHandler(s store.Store) http.HandlerFunc {
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

		if err := hashPassword(user); err != nil {
			service.Error(w, http.StatusInternalServerError, err)
			return
		}

		err := s.User().Create(user)
		if err != nil && err != store.ErrUserAlreadyExist {
			service.Error(w, http.StatusInternalServerError, err)
			return
		}
		if err == store.ErrUserAlreadyExist {
			service.Error(w, http.StatusConflict, err)
			return
		}

		authUser(w, user.Login)
		service.Respond(w, http.StatusOK, "success")
	}
}
