package orders

import (
	"gophermart/internal/app/model"
	"gophermart/internal/app/service"
	"gophermart/internal/app/service/user"
	"gophermart/internal/app/store"
	"io/ioutil"
	"net/http"
)

func OrdersHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderNumber, err := ioutil.ReadAll(r.Body)
		if err != nil {
			service.Error(w, http.StatusBadRequest, err)
			return
		}

		intedNumber, err := bytesToInt(orderNumber)
		if err != nil || !isValid(intedNumber) {
			service.Respond(w, http.StatusUnprocessableEntity, "invalid order code")
			return
		}

		cookie, err := r.Cookie(user.AuthCookieKey)
		if err != nil {
			service.Error(w, http.StatusInternalServerError, err) //here user cookie must exist
		}
		currentOwner := cookie.Value

		oldOwner, err := s.Orders().GetOwnerByNumber(intedNumber)
		if err != nil && err != store.ErrOrderNotExist {
			service.Error(w, http.StatusInternalServerError, err)
			return
		}
		if err == store.ErrOrderNotExist {
			order := &model.Order{
				Number: intedNumber,
				Status: StatusNew,
				Owner:  currentOwner,
			}

			err := s.Orders().Create(order)
			if err != nil {
				service.Error(w, http.StatusInternalServerError, err)
			}

			go calculateBonuces()

			service.Respond(w, http.StatusAccepted, "success")
			return
		}

		if oldOwner == currentOwner {
			service.Respond(w, http.StatusOK, "order already places by this user")
			return
		}

		service.Respond(w, http.StatusConflict, "order already places by another user")
	}
}

func calculateBonuces() {
	//will be implemented when loyalty system will be provided
}
