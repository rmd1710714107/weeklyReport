package main

import (
	"dailyReport/db"
	"dailyReport/routes"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	initSqliteErr := db.InitSqlite()
	if initSqliteErr != nil {
		log.Fatalln(initSqliteErr)
		return
	}
	defer db.Close()
	initGinErr := routes.InitRoutes()
	if initGinErr != nil {
		log.Fatalln(initGinErr)
		return
	}
}
