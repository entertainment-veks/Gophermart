package model

import (
	"time"
)

type Order struct {
	ID          int       `json:"-"`
	Number      string    `json:"number"`
	Status      string    `json:"status"`
	Accrual     *int      `json:"accrual,omitempty"` //nullable
	Owner       string    `json:"-"`
	Uploaded_at time.Time `json:"uploaded_at"`
}
