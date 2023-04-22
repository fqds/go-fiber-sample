package databaseConfig

import (
	"fmt"
	"go-fiber-sample/config"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var Config = viper.New()

func ConnectToDb() *sqlx.DB {
	db, err := ConnectDB(
		config.Config.GetString("db.host"),
		config.Config.GetString("db.port"),
		config.Config.GetString("db.dbname"),
		config.Config.GetString("db.username"),
		config.Config.GetString("db.password"),
		config.Config.GetString("db.sslmode"),
	)
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	log.Println("Database connection success!")

	return db
}

func ConnectDB(host, port, dbname, user, password, sslmode string) (*sqlx.DB, error) {
	databaseURL := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host,
		port,
		dbname,
		user,
		password,
		sslmode,
	)
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
