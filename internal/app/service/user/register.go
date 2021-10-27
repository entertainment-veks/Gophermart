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
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			service.Error(w, http.StatusBadRequest, err)
			return
		}

		if err := validate(user); err != nil {
			service.Error(w, http.StatusBadRequest, err)
		}

		if err := hashPassword(user); err != nil {
			service.Error(w, http.StatusInternalServerError, err)
		}

		switch err := s.User().Create(user); err {
		case store.ErrUserAlreadyExist:
			service.Error(w, http.StatusConflict, err)
		case nil: //no errs
			authUser(w, user.Login)
		default: //any another error
			service.Error(w, http.StatusInternalServerError, err)
		}
	}
}
