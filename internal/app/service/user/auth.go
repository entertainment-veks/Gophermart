package user

import "net/http"

const (
	AuthCookieKey = "gophermart_login"
)

func authUser(w http.ResponseWriter, login string) {
	cookie := &http.Cookie{
		Name:  AuthCookieKey,
		Value: login,
	}
	http.SetCookie(w, cookie)
}

func isUserAuthed(r *http.Request) bool {
	_, err := r.Cookie(AuthCookieKey)
	return err == nil
}
