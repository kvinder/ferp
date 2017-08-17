package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//InitConnection database
func InitConnection(databaseName, databaseURL string) {
	dbConnect, err := sql.Open(databaseName, databaseURL)
	if err != nil {
		log.Fatalf("can not connect database : %v", err)
	}
	db = dbConnect
	fmt.Println("Database connect...")
}

//CloseConnection database
func CloseConnection() {
	db.Close()
	fmt.Println("Database close.")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
