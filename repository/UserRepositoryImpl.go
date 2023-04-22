package repository

import (
	"go-fiber-sample/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepPostgres struct {
	db *sqlx.DB
}

func (b UserRepPostgres) CreateUser(user *entity.User) (int, error) {
	var id int
	if err := b.db.QueryRow("INSERT INTO users (name, encrypted_password) VALUES ($1, $2) RETURNING ID", user.Name, user.EncryptedPassword).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (b UserRepPostgres) GetUserByName(user *entity.User) error {
	if err := b.db.QueryRow("SELECT id, encrypted_password FROM users WHERE name = $1", user.Name).Scan(&user.ID, &user.EncryptedPassword); err != nil {
		return err
	}
	
	return nil
}


func (b UserRepPostgres) GetUserByID(user *entity.User) error {
	if err := b.db.QueryRow("SELECT name FROM users WHERE id = $1", user.ID).Scan(&user.Name); err != nil {
		return err
	}
	
	return nil
}

func NewUserRepPostgres(db *sqlx.DB) *UserRepPostgres {
	return &UserRepPostgres{db: db}
}
