package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models/repositories"
	"github.com/gorilla/mux"
)

// CreateBooking handles the creation of a new booking.
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a booking struct
	var newBooking models.Booking
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Create a Booking instance using the model
	booking := models.NewBooking(0, newBooking.UserID, newBooking.CarID, newBooking.StartTime, newBooking.EndTime)

	// Create a BookingRepository instance with the path to the JSON data file
	bookingRepository := repositories.NewBookingRepositoryJSON("internal/data/bookings.json")

	// Use the repository to create the booking
	err = bookingRepository.CreateBooking(booking)
	if err != nil {
		http.Error(w, "Failed to create booking", http.StatusInternalServerError)
		return
	}

	// Respond with the created booking
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

// GetBooking retrieves a booking by its ID.
func GetBooking(w http.ResponseWriter, r *http.Request) {
	// Extract booking ID from the request
	vars := mux.Vars(r)
	bookingID := vars["id"]

	// Parse the booking ID as an integer (you may want to add error handling)
	id := 0
	fmt.Sscanf(bookingID, "%d", &id)

	// Create a BookingRepository instance with the path to the JSON data file
	bookingRepository := repositories.NewBookingRepositoryJSON("internal/data/bookings.json")

	// Use the repository to get the booking by ID
	booking, err := bookingRepository.GetBookingByID(id)
	if err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	// Respond with the booking details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

// Implement other booking-related handlers as needed (update, delete, list, etc.)
