package model

import "time"

type Operation struct {
	ID           int       `json:"-"`
	Owner        string    `json:"-"`
	Order        string    `json:"order"`
	Amount       int       `json:"sum"`
	Processed_at time.Time `json:"processed_at"`
}
