package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Hinterberger-Thomas/passmo-sev/config"
)

type DB struct {
	client *sql.DB
}

func InitDB() *DB {
	db, err := sql.Open("mysql", createURL())
	if err != nil {
		fmt.Println(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &DB{
		client: db,
	}
}

func createURL() string {
	urlData := config.GetConfMySQL()
	return urlData.UserName + ":" + urlData.DbPass + "@" + urlData.Protocol + "(" + urlData.IP + "" + fmt.Sprint(urlData.Port) + ")/" + urlData.DbName
}
