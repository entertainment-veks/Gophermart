package orders

import (
	"encoding/json"
	"fmt"
	"gophermart/internal/app/handler"
	"gophermart/internal/app/handler/user"
	"gophermart/internal/app/model"
	"gophermart/internal/app/service/orders"
	"gophermart/internal/app/store"
	"io/ioutil"
	"net/http"
)

func OrdersPostHandler(s store.Store, accuralSystemAddress string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input, err := ioutil.ReadAll(r.Body)
		if err != nil {
			handler.Error(w, http.StatusBadRequest, err)
			return
		}

		orderNumber := string(input)

		if err != nil || !orders.IsValid(orderNumber) {
			handler.Respond(w, http.StatusUnprocessableEntity, "invalid order code")
			return
		}

		currentOwner, err := user.GetLogin(r)
		if err != nil {
			handler.Error(w, http.StatusInternalServerError, err)
			return
		}

		oldOwner, err := s.Orders().GetOwnerByNumber(orderNumber)
		if err != nil && err != store.ErrOrderNotExist {
			handler.Error(w, http.StatusInternalServerError, err)
			return
		}
		if err == store.ErrOrderNotExist {
			order := &model.Order{
				Number: orderNumber,
				Status: orders.StatusNew,
				Owner:  currentOwner,
			}

			err := s.Orders().Create(order)
			if err != nil {
				handler.Error(w, http.StatusInternalServerError, err)
				return
			}

			go calculateBonuces(s, accuralSystemAddress, orderNumber)

			handler.Respond(w, http.StatusAccepted, "success")
			return
		}

		if oldOwner == currentOwner {
			handler.Respond(w, http.StatusOK, "order already places by this user")
			return
		}

		handler.Respond(w, http.StatusConflict, "order already places by another user")
	}
}

func calculateBonuces(s store.Store, accuralSystemAddress string, orderNumber string) {
	link := fmt.Sprintf("%s/api/orders/%s", accuralSystemAddress, orderNumber)
	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	var order *model.Order
	if err := json.NewDecoder(resp.Body).Decode(&order); err != nil {
		return
	}

	if err := s.Orders().UpdateStatus(order); err != nil {
		return
	}
}
