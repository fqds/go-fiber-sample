package entity

type User struct {
	ID                int    `json:"-" db:"id"`
	Name              string `json:"-" db:"name"`
	EncryptedPassword string `json:"-" db:"encrypted_password"`
}
