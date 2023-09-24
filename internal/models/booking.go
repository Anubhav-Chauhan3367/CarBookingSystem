package models

import (
	"time"
)

// Booking represents a car booking made by a user.
type Booking struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    CarID     int       `json:"car_id"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    // Add more fields as needed
}
type BookingAvailability struct {
    ID int    `json:"id"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
}

// NewBooking creates a new Booking instance.
func NewBooking(id, userID, carID int, startTime, endTime time.Time) *Booking {
    return &Booking{
        ID:        id,
        UserID:    userID,
        CarID:     carID,
        StartTime: startTime,
        EndTime:   endTime,
    }
}
