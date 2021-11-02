package user

import (
	"gophermart/internal/app/service"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isUserAuthed(r) {
			next.ServeHTTP(w, r)
			return
		}
		service.Respond(w, http.StatusUnauthorized, "user is unauthorized")
	})
}
