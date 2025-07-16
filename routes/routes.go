package routes

import (
	"mini-trading-platform/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}).Methods("GET")

	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/stocks", controllers.GetStocks).Methods("GET")
	router.HandleFunc("/order", controllers.PlaceOrder).Methods("POST")
	router.HandleFunc("/portfolio", controllers.GetPortfolio).Methods("GET")
	router.HandleFunc("/orders", controllers.GetOrders).Methods("GET")

	return router
}
