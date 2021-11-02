package model

import (
	"database/sql"
	"time"
)

type Order struct {
	ID            int           `json:"-"`
	Number        int           `json:"number"`
	Status        string        `json:"status"`
	Accrual       int           `json:"accrual"`
	AccrualFromDB sql.NullInt16 `json:"-"`
	Owner         string        `json:"-"`
	Uploaded_at   time.Time     `json:"uploaded_at"`
}
