package user

import (
	"gophermart/internal/app/handler"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isUserAuthed(r) {
			next.ServeHTTP(w, r)
			return
		}
		handler.Respond(w, http.StatusUnauthorized, "user is unauthorized")
	})
}
