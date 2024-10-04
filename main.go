package main

import (

	"log"
	"beverages-booking/config"
	"beverages-booking/server"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	log.Println("Starting beverage_booking App")

	log.Println("Initializig configuration")
	config := config.InitConfig("beverage_booking")

	log.Println("Initializig database")
	dbHandler := server.InitDatabase(config)

	log.Println("Initializig HTTP sever")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()
}
