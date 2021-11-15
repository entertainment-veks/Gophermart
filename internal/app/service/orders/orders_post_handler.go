package orders

import (
	"gophermart/internal/app/model"
	"gophermart/internal/app/service"
	"gophermart/internal/app/service/user"
	"gophermart/internal/app/store"
	"io/ioutil"
	"net/http"
)

func OrdersPostHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input, err := ioutil.ReadAll(r.Body)
		if err != nil {
			service.Error(w, http.StatusBadRequest, err)
			return
		}

		orderNumber := string(input)

		if err != nil || !IsValid(orderNumber) {
			service.Respond(w, http.StatusUnprocessableEntity, "invalid order code")
			return
		}

		currentOwner, err := user.GetLogin(r)
		if err != nil {
			service.Error(w, http.StatusInternalServerError, err)
			return
		}

		oldOwner, err := s.Orders().GetOwnerByNumber(orderNumber)
		if err != nil && err != store.ErrOrderNotExist {
			service.Error(w, http.StatusInternalServerError, err)
			return
		}
		if err == store.ErrOrderNotExist {
			order := &model.Order{
				Number: orderNumber,
				Status: StatusNew,
				Owner:  currentOwner,
			}

			err := s.Orders().Create(order)
			if err != nil {
				service.Error(w, http.StatusInternalServerError, err)
				return
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
