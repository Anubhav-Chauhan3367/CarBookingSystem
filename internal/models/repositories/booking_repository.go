package repositories

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models"
)

// BookingRepositoryJSON is an implementation of BookingRepository using JSON files.
type BookingRepositoryJSON struct {
	filePath string
	mu       sync.Mutex // Mutex for concurrent access
}

func NewBookingRepositoryJSON(filePath string) *BookingRepositoryJSON {
	return &BookingRepositoryJSON{filePath: filePath}
}

func (r *BookingRepositoryJSON) GetBookingByID(id int) (*models.Booking, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var bookings []*models.Booking
	if err := json.Unmarshal(data, &bookings); err != nil {
		return nil, err
	}

	for _, booking := range bookings {
		if booking.ID == id {
			return booking, nil
		}
	}

	return nil, fmt.Errorf("booking with ID %d not found", id)
}

func (r *BookingRepositoryJSON) GetAllBookings() ([]*models.Booking, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var bookings []*models.Booking
	if err := json.Unmarshal(data, &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepositoryJSON) CreateBooking(booking *models.Booking) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	var bookings []*models.Booking
	if err := json.Unmarshal(data, &bookings); err != nil {
		return err
	}

	// Assign a unique ID to the new booking
	if len(bookings) == 0 {
		booking.ID = 1
	} else {
		lastBooking := bookings[len(bookings)-1]
		booking.ID = lastBooking.ID + 1
	}

	bookings = append(bookings, booking)

	newData, err := json.Marshal(bookings)
	if err != nil {
		return err
	}

	if err := os.WriteFile(r.filePath, newData, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// Implement UpdateBooking and DeleteBooking methods similarly
