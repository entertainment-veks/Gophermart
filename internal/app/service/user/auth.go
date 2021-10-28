package user

import "net/http"

const (
	authCookieKey = "gophermart-login"
)

func authUser(w http.ResponseWriter, login string) {
	cookie := &http.Cookie{
		Name:  authCookieKey,
		Value: login,
	}
	http.SetCookie(w, cookie)
}

func isUserAuthed() {} //todo
