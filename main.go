package main

import (
	"beverages_booking/config"
	"beverages_booking/server"
)

func main() {
	config.InitDB()

	server.Start()
}