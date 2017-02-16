package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
)

type configuration struct {
	DBHost, DBUser, DBPwd, Database string
}

func CreateDbSession() (db *sql.DB, err error) {
	var AppConfig configuration

	AppConfig = getDBConfig()

	setup := fmt.Sprintf("%s:%s@tcp(%s)/%s", AppConfig.DBUser, AppConfig.DBPwd, AppConfig.DBHost, AppConfig.Database)
	db, err = sql.Open("mysql", setup)
	if err != nil {
		return nil,err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return
}

func getDBConfig() (AppConfig configuration) {

	dbHost := os.Getenv("SPORTS_DATABASE_HOST")
	dbUser := os.Getenv("SPORTS_DATABASE_USER")
	dbPwd := os.Getenv("SPORTS_DATABASE_PWD")
	db := os.Getenv("DATABASE_SPORTS")

	//fmt.Println("keys:", dbHost, dbUser, dbPwd)

	AppConfig = configuration{
		DBHost: dbHost,
		DBUser: dbUser,
		DBPwd: dbPwd,
		Database: db,
	}
	return
}
