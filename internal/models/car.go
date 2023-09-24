package models

// Car represents a car available for rental.
type Car struct {
    ID          int    `json:"id"`
    Brand       string `json:"brand"`
    Model       string `json:"model"`
    Year        int    `json:"year"`
    Description string `json:"description"`
    // Add more fields as needed
}

// CarWithAvailability represents a car along with its availability data.
type CarWithAvailability struct {
    Car         *Car          `json:"car"`
    Availability []Availability `json:"availability"`
}

// NewCar creates a new Car instance.
func NewCar(id int, brand, model string, year int, description string) *Car {
    return &Car{
        ID:          id,
        Brand:       brand,
        Model:       model,
        Year:        year,
        Description: description,
    }
}
