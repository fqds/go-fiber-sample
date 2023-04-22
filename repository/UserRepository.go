package repository

import (
	"go-fiber-sample/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user *entity.User) (int, error)
	GetUserByName(user *entity.User) error
	GetUserByID(user *entity.User) error
}

type UserRep struct {
	UserRepository
}

func NewUserRep(db *sqlx.DB) *UserRep {
	return &UserRep{UserRepository: NewUserRepPostgres(db)}
}
