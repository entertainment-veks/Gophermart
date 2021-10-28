package model

type User struct {
	ID       int
	Login    string `json:"login"`
	Password string `json:"password"`
}
