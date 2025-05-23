package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/terrnit/rebound/backend/docs" // This will be generated
	pizzahandler "github.com/terrnit/rebound/backend/internal/delivery/http/pizza"
	pizzarepo "github.com/terrnit/rebound/backend/internal/repository/pizza"
	pizzausecase "github.com/terrnit/rebound/backend/internal/usecase/pizza"
)

// @title Pizza CRUD API
// @version 1.0
// @description A simple Pizza CRUD API service
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize repository
	repo := pizzarepo.NewMemoryRepository()

	// Initialize use case
	useCase := pizzausecase.NewUseCase(repo)

	// Initialize handler
	handler := pizzahandler.NewHandler(useCase)

	// Setup router
	router := mux.NewRouter()

	// Swagger documentation
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// Register routes
	router.HandleFunc("/pizzas", handler.Create).Methods(http.MethodPost)
	router.HandleFunc("/pizzas/{id:[0-9]+}", handler.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/pizzas/{id:[0-9]+}", handler.Update).Methods(http.MethodPut)
	router.HandleFunc("/pizzas/{id:[0-9]+}", handler.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/pizzas", handler.List).Methods(http.MethodGet)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
