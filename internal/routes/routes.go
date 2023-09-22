package routes

import (
	"github.com/gorilla/mux"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/handlers"
)

// Set up routes for your API
func SetupRoutes(r *mux.Router) {
    // User-related routes
    r.HandleFunc("/api/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/api/users/{id}", handlers.GetUser).Methods("GET")
    // Add more user-related routes as needed

    // Car-related routes
    r.HandleFunc("/api/cars", handlers.CreateCar).Methods("POST")
    r.HandleFunc("/api/cars/{id}", handlers.GetCar).Methods("GET")
    // Add more car-related routes as needed

    // Booking-related routes
    r.HandleFunc("/api/bookings", handlers.CreateBooking).Methods("POST")
    r.HandleFunc("/api/bookings/{id}", handlers.GetBooking).Methods("GET")
    // Add more booking-related routes as needed

    // Define other routes for your application
}
