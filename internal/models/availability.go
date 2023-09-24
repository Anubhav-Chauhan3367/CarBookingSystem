// models/availability.go

package models

import "time"

// Availability represents the time slot when a car is available.
type Availability struct {
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
}
