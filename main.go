package main

import (
	"fmt"
	"log"
	"net/http"

	"mini-trading-platform/database"
	"mini-trading-platform/routes"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Mini Trading Platform Server...")
	database.Connect()

	router := routes.SetupRoutes()

	// 🟡 Allow frontend to access backend (CORS fix)
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(corsMiddleware)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
