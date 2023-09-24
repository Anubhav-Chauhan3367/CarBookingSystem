package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"sort"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models/repositories"
	"github.com/gorilla/mux"
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

func GetCarsWithAvailabilityData(w http.ResponseWriter, r *http.Request) {
	// Create a CarRepository instance with the path to the JSON data file
	carRepository := repositories.NewCarRepositoryJSON("internal/data/cars.json")

	// Create a BookingRepository instance with the path to the JSON data file
	bookingRepository := repositories.NewBookingRepositoryJSON("internal/data/booking.json")

	// Get all cars from the car repository
	cars, err := carRepository.GetAllCars()
	if err != nil {
		http.Error(w, "Failed to fetch cars", http.StatusInternalServerError)
		return
	}

	// Get all bookings from the booking repository
	bookings, err := bookingRepository.GetAllBookings()
	if err != nil {
		http.Error(w, "Failed to fetch bookings", http.StatusInternalServerError)
		return
	}

	// Create a map to store car availability data
	carAvailability := make(map[int][]models.BookingAvailability)

	// Initialize car availability based on car IDs
	for _, car := range cars {
		carAvailability[car.ID] = []models.BookingAvailability{}
	}

	// Populate car availability data based on bookings
	for _, booking := range bookings {
		if booking.CarID > 0 {
			// Append booking availability data
			bookingAvailability := models.BookingAvailability{
				ID:        booking.ID,
				StartTime: booking.StartTime,
				EndTime:   booking.EndTime,
			}

			carAvailability[booking.CarID] = append(carAvailability[booking.CarID], bookingAvailability)
		}
	}

	// Create a slice to store cars with availability data
	carsWithAvailability := make([]models.CarWithAvailability, 0)

	// Populate cars with availability data
	for _, car := range cars {
		// Calculate car availability based on bookings
		availability := calculateCarAvailability(carAvailability[car.ID])
		carWithAvailability := models.CarWithAvailability{
			Car:         car,
			Availability: availability,
		}

		carsWithAvailability = append(carsWithAvailability, carWithAvailability)
	}

	// Respond with the cars with availability data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carsWithAvailability)
}

func calculateCarAvailability(bookings []models.BookingAvailability) []models.Availability {
	// Sort the bookings by start time
	sort.Slice(bookings, func(i, j int) bool {
		return bookings[i].StartTime.Before(bookings[j].StartTime)
	})

	availability := []models.Availability{}

	// Calculate availability based on sorted bookings
	if len(bookings) > 0 {
		start := bookings[0].EndTime
		for _, booking := range bookings[1:] {
			end := booking.StartTime
			availability = append(availability, models.Availability{
				StartTime: start,
				EndTime:   end,
			})
			start = booking.EndTime
		}
		// Handle the last available time slot
		availability = append(availability, models.Availability{
			StartTime: start,
			EndTime:   bookings[len(bookings)-1].EndTime,
		})
	} else {
		// If no bookings exist, the car is available all the time
		availability = append(availability, models.Availability{
			StartTime: time.Time{},
			EndTime:   time.Time{},
		})
	}

	return availability
}


// Implement other car-related handlers as needed (update, delete, list, etc.)
