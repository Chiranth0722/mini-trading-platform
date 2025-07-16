package main

import (
	"fmt"
	"log"
	"net/http"

	"mini-trading-platform/database"
	"mini-trading-platform/routes"
)

func main() {
	fmt.Println("Starting Mini Trading Platform Server...")

	// Connect to PostgreSQL
	database.Connect()

	// Setup routes from routes.go
	router := routes.SetupRoutes()

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
