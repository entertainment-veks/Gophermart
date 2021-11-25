package user

import (
	"fmt"
	"gophermart/internal/app/handler"
	"gophermart/internal/app/model"
	"gophermart/internal/app/store"
	"log"
)

func Register(s store.Store, u *model.User) error {
	if err := validate(u); err != nil {
		return fmt.Errorf("%w: %v", handler.ErrBadRequest, err)
	}

	if err := hashPassword(u); err != nil {
		log.Println(err)
		return fmt.Errorf("%w: %v", handler.ErrInternalServer, err)
	}

	err := s.User().Create(u)
	if err != nil && err != store.ErrUserAlreadyExist {
		log.Println(err)
		return fmt.Errorf("%w: %v", handler.ErrInternalServer, err)
	}
	if err == store.ErrUserAlreadyExist {
		return fmt.Errorf("%w: %v", handler.ErrConflict, err)
	}

	return nil
}

func Login(s store.Store, u *model.User) (bool, error) {
	if err := validate(u); err != nil {
		return false, fmt.Errorf("%w: %v", handler.ErrBadRequest, err)
	}

	foundUser, err := s.User().GetByLogin(u.Login)
	if err != nil && err != store.ErrUserNotFound {
		return false, fmt.Errorf("%w: %v", handler.ErrInternalServer, err)
	}
	if err == store.ErrUserNotFound {
		return false, fmt.Errorf("%w: %v", handler.ErrUnauthorized, err)
	}

	if err := hashPassword(u); err != nil {
		return false, fmt.Errorf("%w: %v", handler.ErrInternalServer, err)
	}

	return equal(u, foundUser), nil
}
