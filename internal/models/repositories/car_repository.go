package repositories

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models"
)

// CarRepositoryJSON is an implementation of CarRepository using JSON files.
type CarRepositoryJSON struct {
	filePath string
	mu       sync.Mutex // Mutex for concurrent access
}

func NewCarRepositoryJSON(filePath string) *CarRepositoryJSON {
	return &CarRepositoryJSON{filePath: filePath}
}

func (r *CarRepositoryJSON) GetCarByID(id int) (*models.Car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var cars []*models.Car
	if err := json.Unmarshal(data, &cars); err != nil {
		return nil, err
	}

	for _, car := range cars {
		if car.ID == id {
			return car, nil
		}
	}

	return nil, fmt.Errorf("car with ID %d not found", id)
}

func (r *CarRepositoryJSON) GetAllCars() ([]*models.Car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var cars []*models.Car
	if err := json.Unmarshal(data, &cars); err != nil {
		return nil, err
	}

	return cars, nil
}

func (r *CarRepositoryJSON) CreateCar(car *models.Car) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	var cars []*models.Car
	if err := json.Unmarshal(data, &cars); err != nil {
		return err
	}

	// Assign a unique ID to the new car
	if len(cars) == 0 {
		car.ID = 1
	} else {
		lastCar := cars[len(cars)-1]
		car.ID = lastCar.ID + 1
	}

	cars = append(cars, car)

	newData, err := json.Marshal(cars)
	if err != nil {
		return err
	}

	if err := os.WriteFile(r.filePath, newData, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// Implement UpdateCar and DeleteCar methods similarly
