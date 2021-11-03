package user

import "net/http"

const (
	authCookieKey = "gophermart_login"
)

func authUser(w http.ResponseWriter, login string) {
	cookie := &http.Cookie{
		Name:  authCookieKey,
		Value: login,
	}
	http.SetCookie(w, cookie)
}

func isUserAuthed(r *http.Request) bool {
	_, err := r.Cookie(authCookieKey)
	return err == nil
}

func GetLogin(r *http.Request) (string, error) {
	cookie, err := r.Cookie(authCookieKey)
	return cookie.Value, err
}
