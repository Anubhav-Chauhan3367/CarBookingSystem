package main

import (
	"fmt"
	"net/http"

	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    routes.SetupRoutes(r)

    // Configure CORS
    cors := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Replace with your frontend origin
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
        handlers.AllowedHeaders([]string{"Content-Type"}),
        handlers.AllowCredentials(),
    )(r)

    // Create an HTTP server with your router
    srv := &http.Server{
        Handler: cors,
        Addr:    "localhost:8080", // Add "http://" here
    }

    // Start the server
    fmt.Println("Server Started on localhost:8080")
    if err := srv.ListenAndServe(); err != nil {
        panic(err)
    }
}
