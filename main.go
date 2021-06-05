package main

import (
	"go-lang-rest-api/database"
	"go-lang-rest-api/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()

}
