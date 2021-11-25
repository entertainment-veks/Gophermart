package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrUnauthorized   = errors.New("user is unauthorized")
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
	ErrConflict       = errors.New("conflict")
)

func Error(w http.ResponseWriter, code int, err error) {
	Respond(w, code, err.Error())
}

func Respond(w http.ResponseWriter, code int, data string) {
	w.WriteHeader(code)
	if len(data) != 0 {
		w.Write([]byte(data))
	}
}

func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}
	Respond(w, statusCode, "")
}
