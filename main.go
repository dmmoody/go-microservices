package main

import (
	"githab.com/dmmoody/go-microservices/internal/database"
	"githab.com/dmmoody/go-microservices/internal/server"
	"log"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to create database client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
