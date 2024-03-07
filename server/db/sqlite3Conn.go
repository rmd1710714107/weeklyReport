package db

import (
	"database/sql"
	"fmt"
	"log"
)

var sqlite3 *sql.DB

func InitSqlite() error {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalln(err)
		return err
	}
	sqlite3 = db
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalln(pingErr)
		return pingErr
	}
	fmt.Println("connect success")
	return nil

}
func Close() {
	sqlite3.Close()
}
