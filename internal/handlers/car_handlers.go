package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models/repositories"
)

// CreateCar handles the creation of a new car.
func CreateCar(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a car struct
	var newCar models.Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Create a new car instance using the model
	car := models.NewCar(0, newCar.Brand, newCar.Model, newCar.Year, newCar.Description)

	// Create a CarRepository instance with the path to the JSON data file
	carRepository := repositories.NewCarRepositoryJSON("internal/data/cars.json")

	// Use the repository to create the car
	
	err = carRepository.CreateCar(car)
	if err != nil {
		http.Error(w, "Failed to create car", http.StatusInternalServerError)
		return
	}

	// Respond with the created car
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

// GetCar retrieves a car by its ID.
func GetCar(w http.ResponseWriter, r *http.Request) {
	// Extract car ID from the request
	vars := mux.Vars(r)
	carID := vars["id"]

	// Parse the car ID as an integer (you may want to add error handling)
	id := 0
	fmt.Sscanf(carID, "%d", &id)

	// Create a CarRepository instance with the path to the JSON data file
	carRepository := repositories.NewCarRepositoryJSON("internal/data/cars.json")

	// Use the repository to get the car by ID
	car, err := carRepository.GetCarByID(id)
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	// Respond with the car details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

// Implement other car-related handlers as needed (update, delete, list, etc.)
