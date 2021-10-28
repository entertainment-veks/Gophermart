package user

import (
	"encoding/hex"
	"gophermart/internal/app/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func validate(u *model.User) error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Login, validation.Required),
		validation.Field(&u.Password, validation.Required),
	)
}

func hashPassword(u *model.User) error {
	u.Password = hex.EncodeToString([]byte(u.Password))
	return nil
}

func equal(f *model.User, s *model.User) bool {
	return f.Login == s.Login && f.Password == s.Password
}
