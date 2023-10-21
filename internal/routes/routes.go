package routes

import (
	"net/http"

	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/handlers"
	"github.com/gorilla/mux"
)

// AuthMiddleware is a middleware that checks if the user is authenticated.

// Set up routes for your API
func SetupRoutes(r *mux.Router) {
    //Home Page/Default Page routes
    r.Handle("/", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetCarsWithAvailabilityData))).Methods("GET")

    // User-related routes
    r.Handle("/api/users", handlers.AuthMiddleware(http.HandlerFunc(handlers.CreateUser))).Methods("POST")
    r.Handle("/api/users/{id}", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetUser))).Methods("GET")
    // Add more user-related routes as needed


    //Authentication-related routes
    r.HandleFunc("/check-auth", handlers.AuthenticateUser).Methods("GET")
    r.HandleFunc("/register", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
    r.HandleFunc("/logout", handlers.LogoutUser).Methods("POST")   
    // Add more authentication-related routes as needed


    // Car-related routes
    r.Handle("/api/allcars", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetAllCars))).Methods("GET")
    r.Handle("/api/cars", handlers.AuthMiddleware(http.HandlerFunc(handlers.CreateCar))).Methods("POST")
    r.Handle("/api/cars/{id}", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetCar))).Methods("GET")
    // Add more car-related routes as needed 


    // Booking-related routes
    r.Handle("/api/bookings", handlers.AuthMiddleware(http.HandlerFunc(handlers.CreateBooking))).Methods("POST")
    r.Handle("/api/bookings/{id}", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetBooking))).Methods("GET")
    // Add more booking-related routes as needed

    // Define other routes for your application
}
