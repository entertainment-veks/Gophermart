package store

import "errors"

var (
	ErrUserAlreadyExist   = errors.New("user already exist")
	ErrUserNotFound       = errors.New("user not found")
	ErrOrderNotExist      = errors.New("order does not exist")
	ErrOrdersNotFound     = errors.New("orders not found")
	ErrOperationsNotFound = errors.New("operations not found")
)
