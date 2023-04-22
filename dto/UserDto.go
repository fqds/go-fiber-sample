package dto

import (
	"go-fiber-sample/exception"

	"golang.org/x/crypto/bcrypt"
)

type UserDto struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"-" db:"name"`
	Password string `json:"-" db:"password"`
}

func (u *UserDto) GetEncryptedPassword() (string, error) {
	if len(u.Password) == 0 {
		return "", exception.ShortPassword{}
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}

	return string(encryptedPassword), nil
}
func (u *UserDto) ComparePassword(encryptedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(u.Password)) == nil
}
