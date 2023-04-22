package service

import (
	"go-fiber-sample/config"
	"go-fiber-sample/databaseConfig"
	"go-fiber-sample/dto"
	"go-fiber-sample/entity"
	"go-fiber-sample/exception"
	"go-fiber-sample/repository"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateUser(user *dto.UserDto) error {
	encryptedPassword, err := user.GetEncryptedPassword()
	if err != nil {
		return err
	}

	userToCreate := &entity.User{
		Name:              user.Name,
		EncryptedPassword: encryptedPassword,
	}
	user.ID, err = repository.NewUserRep(databaseConfig.ConnectToDb()).CreateUser(userToCreate)
	if err != nil {
		return exception.NotCreatedObject{}
	}
	return nil
}

func CreateSession(user *dto.UserDto) (string, error) {

	userWithData := &entity.User{
		Name: user.Name,
	}

	if err := repository.NewUserRep(databaseConfig.ConnectToDb()).GetUserByName(userWithData); err != nil {
		return "", err
	}

	if !user.ComparePassword(userWithData.EncryptedPassword) {
		return "", exception.NotAuthenticated{}
	}

	expirationTime := time.Now().Add(time.Second * 900)
	claims := &Claims{
		UserID: userWithData.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Config.GetString("auth-secret")))
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return signedToken, nil
}

func ApproveSession(authToken string) (*entity.User, error) {

	claims := &Claims{}
	_, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.GetString("auth-secret")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, exception.NotAuthenticated{}
	}
	user := &entity.User{
		ID: claims.UserID,
	}
	
	if err := repository.NewUserRep(databaseConfig.ConnectToDb()).GetUserByID(user); err != nil {
		return nil, err
	}

	return user, nil
}
