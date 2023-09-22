package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models/repositories"
	"github.com/gorilla/mux"
)

// CreateUser handles the creation of a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a user struct
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Create a new user instance using the model
	user := models.NewUser(0, newUser.Username, newUser.Email, newUser.Password)

	// Create a UserRepository instance with the path to the JSON data file
	userRepository := repositories.NewUserRepositoryJSON("internal/data/users.json")
	// log.Println(
	// 	"UserRepositoryJSON: ", userRepository,
	// 	"User: ", user,
	// )
	// Use the repository to create the user
	err = userRepository.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Respond with the created user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser retrieves a user by its ID.
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the request
	vars := mux.Vars(r)
	userID := vars["id"]
	
	// Parse the user ID as an integer (you may want to add error handling)
	id := 0
	fmt.Sscanf(userID, "%d", &id)

	// Create a UserRepository instance with the path to the JSON data file
	userRepository := repositories.NewUserRepositoryJSON("internal/data/users.json")

	// Use the repository to get the user by ID
	user, err := userRepository.GetUserByID(id)
	log.Println(
		"UserRepositoryJSON: ", userRepository,
		id,
	)
	
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Respond with the user details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Implement other user-related handlers as needed (update, delete, list, etc.)
