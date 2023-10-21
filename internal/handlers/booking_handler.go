package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models/repositories"
	"github.com/gorilla/mux"
)

// CreateBooking handles the creation of a new booking.
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a booking struct
	var newBooking models.Booking
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	fmt.Println(err, "check")
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Create a Booking instance using the model
	booking := models.NewBooking(0, newBooking.UserID, newBooking.CarID, newBooking.StartTime, newBooking.EndTime)
	fmt.Println(booking, "check2")
	// Create a BookingRepository instance with the path to the JSON data file
	bookingRepository := repositories.NewBookingRepositoryJSON("internal/data/booking.json")

	// Create a CarRepository instance with the path to the JSON data file
	carRepository := repositories.NewCarRepositoryJSON("internal/data/cars.json")

	// Get the car associated with the booking
	car, err := carRepository.GetCarByID(booking.CarID)
	// fmt.Println(err, "check3", car, "check4", bookingRepository, "check5")
	if err != nil {
		http.Error(w, "Failed to fetch car", http.StatusInternalServerError)
		return
	}

	// Check if the car is available for the requested time slot
	fmt.Println(booking.CarID, " ", booking.UserID, " ", booking.StartTime, " ", booking.EndTime, " ", "check6")
	if !IsCarAvailableForBooking(*car, booking.StartTime, booking.EndTime) {
		http.Error(w, "Car is not available for the requested time slot", http.StatusBadRequest)
		return
	}

	// Use the repository to create the booking
	err = bookingRepository.CreateBooking(booking)
	
	fmt.Println(err, "check")
	if err != nil {
		http.Error(w, "Failed to create booking", http.StatusInternalServerError)
		return
	}

	// Respond with the created booking
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func IsCarAvailableForBooking(car models.Car, startTime time.Time, endTime time.Time) bool {
	// Create a BookingRepository instance with the path to the JSON data file
	bookingRepository := repositories.NewBookingRepositoryJSON("internal/data/booking.json")

	// Get all bookings for the car
	bookings, err := bookingRepository.GetBookingByID(car.ID)
	fmt.Println(bookings, "check7")
	if err != nil {
		// Handle the error (e.g., log or return false)
		return false
	}

	// Iterate through existing bookings
	for _, booking := range bookings {
		// Check for overlaps

		if booking.StartTime.Before(endTime) && booking.EndTime.After(startTime) {
			// There is an overlap, so the car is not available
			fmt.Println("check8")
			return false
		}
	}

	// No overlaps found, so the car is available
	return true
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
	bookingRepository := repositories.NewBookingRepositoryJSON("internal/data/booking.json")

	// Use the repository to get the booking by ID
	bookings, err := bookingRepository.GetBookingByID(id)
	if err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	// Respond with the booking details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// Implement other booking-related handlers as needed (update, delete, list, etc.)
