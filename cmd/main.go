package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/routes"
)

func main() {
    // Initialize a new Gorilla Mux router
    r := mux.NewRouter()

    // Setup routes by calling the SetupRoutes function from the routes package
    routes.SetupRoutes(r)

    // Start the web server
    port := 8080
    fmt.Printf("Server started on :%d\n", port)
    http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
