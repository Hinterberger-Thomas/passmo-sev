package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Hinterberger-Thomas/passmo-sev/config"
	"github.com/Hinterberger-Thomas/passmo-sev/graph/model"
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

func (db *DB) InsNewAcc(data model.UserD) (bool, error) {
	insStmt := "INSERT INTO USER (email, password) VALUES (?,?);"
	stmt, err := db.client.Query(insStmt, data.Email, data.Password)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	return true, err
}

func (db *DB) UpdAccData(data model.AccountD) (bool, error) {
	updStmt := "UPDATE data SET password = ?, data = ?, username = ? WHERE name = ?; "
	stmt, err := db.client.Query(updStmt, data.Password, data.Data, data.Usern, data.Name)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	return true, err
}
