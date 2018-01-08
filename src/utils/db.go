package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var DB *sql.DB

type config struct {
	username string
	password string
	ip       string
	port     string
	database string
	charset  string
}

func init() {

	cf := config{"root", "", "127.0.0.1", "3306", "blog", "utf8"}

	dbUrl := cf.username + ":" + cf.password + "@tcp(" + cf.ip + ":" + cf.port + ")/" + cf.database + "?charset=" + cf.charset

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}
