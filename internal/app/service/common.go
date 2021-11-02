package service

import (
	"encoding/json"
	"net/http"
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
