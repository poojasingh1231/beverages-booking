package config

import (
	"database/sql"
	"fmt"
	"log"
	// "os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", "experiment", "experiment", "localhost", "beverages_booking")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println("Database connected successfully.")
}
