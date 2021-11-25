package user

import (
	"gophermart/internal/app/handler"
	"math/rand"
	"net/http"
	"time"
)

const (
	authCookieKey = "gophermart_session"
)

type Session struct {
	Login     string
	ExpiredAt time.Time
}

var sessions map[string]Session = make(map[string]Session)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func authUser(w http.ResponseWriter, login string) {
	sessionID := generateSessionID()
	sessions[sessionID] = Session{
		Login:     login,
		ExpiredAt: time.Now().Add(time.Minute * 30), //so it will expire after a half of hour
	}

	cookie := &http.Cookie{
		Name:  authCookieKey,
		Value: sessionID,
	}
	http.SetCookie(w, cookie)
}

func isUserAuthed(r *http.Request) bool {
	cookie, err := r.Cookie(authCookieKey)
	if err != nil {
		return false
	}
	if sessions[cookie.Value].ExpiredAt.Before(time.Now()) {
		return false
	}
	return true
}

func GetLogin(r *http.Request) (string, error) {
	if isUserAuthed(r) {
		cookie, _ := r.Cookie(authCookieKey)
		return sessions[cookie.Value].Login, nil
	}
	return "", handler.ErrUnauthorized
}

func generateSessionID() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 30)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
