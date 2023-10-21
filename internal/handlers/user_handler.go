package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models/repositories"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check the user's session to see if they are authenticated
        session, _ := store.Get(r, "user-session")
        userID, ok := session.Values["user_id"].(int)
		fmt.Println(userID, "this")
        if !ok {
			fmt.Println("this")
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}


// CreateUser handles the creation of a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a user struct
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Check if the username already exists in the repository
	userRepository := repositories.NewUserRepositoryJSON("internal/data/users.json")
	existingUser, err := userRepository.GetUserByUsername(newUser.Username)
	if err == nil {
		fmt.Println(existingUser)
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Create a new user in the repository
	err = userRepository.CreateUser(&newUser)
	if err != nil {
		fmt.Println(err, "check")
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Respond with a success message or user data in JSON format
	userResponse := struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}{
		ID:       newUser.ID,
		Username: newUser.Username,
	}

	response, err := json.Marshal(userResponse)
	if err != nil {
		fmt.Println("here")
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	// fmt.Println(loginUser, " | ", err)

	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Create a UserRepository instance with the path to the JSON data file
	userRepository := repositories.NewUserRepositoryJSON("internal/data/users.json")
	// fmt.Println(userRepository)
	// Use the repository to retrieve the user by username
	storedUser, err := userRepository.GetUserByUsername(loginUser.Username)
	// fmt.Println(storedUser, " | ", err)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare the provided password with the stored password (you should use a secure password hashing library in a real system)
	if loginUser.Password != storedUser.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create a session when the user logs in
	session, _ := store.Get(r, "user-session")
	// fmt.Println(storedUser.ID)
	session.Values["user_id"] = storedUser.ID
	session.Save(r, w)
	// fmt.Println(session.Values, "this")
	// Respond with user data in JSON format
	userResponse := struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}{
		ID:       storedUser.ID,
		Username: storedUser.Username,
	}

	// Encode and send the JSON response
	response, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
	// fmt.Println(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}


func LogoutUser(w http.ResponseWriter, r *http.Request) {
	// Clear the user's session to log them out
	session, _ := store.Get(r, "user-session")
	session.Options = &sessions.Options{MaxAge: -1}
	session.Save(r, w)
	fmt.Println(session.Options, "this")
	w.WriteHeader(http.StatusOK)
}


// GetUser retrieves a user by its ID.
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the request
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Println(userID)
	// Parse the user ID as an integer (you may want to add error handling)
	id := 0
	// fmt.Sscanf(userID, "%d", &id)

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

// AuthenticateUser is a handler for checking if a user is authenticated.
// AuthenticateUser is a handler for checking if a user is authenticated.
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
    if userIsAuthenticated(r) {
        // User is authenticated, respond with user data in JSON format.
        session, _ := store.Get(r, "user-session")
        userID, _ := session.Values["user_id"].(int)

        // Create a UserRepository instance with the path to the JSON data file
        userRepository := repositories.NewUserRepositoryJSON("internal/data/users.json")

        // Use the repository to retrieve the user by user ID
        user, err := userRepository.GetUserByID(userID)
        if err != nil {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }

        // Construct the user response structure
        userResponse := struct {
            ID       int    `json:"id"`
            Username string `json:"username"`
        }{
            ID:       user.ID,
            Username: user.Username,
        }

        // Encode and send the JSON response
        response, err := json.Marshal(userResponse)
        if err != nil {
            http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(response)
    } else {
        // User is not authenticated, respond with an unauthorized status.
        http.Error(w, "Not authenticated", http.StatusUnauthorized)
    }
}



func userIsAuthenticated(r *http.Request) bool {
	session, err := store.Get(r, "user-session")
	if err != nil {
		// fmt.Println(err, "here")
		return false
	}


	if userID, ok := session.Values["user_id"]; ok {
		// fmt.Println(userID)
		return userID != nil
	}

	// Return false by default if no authentication data is found.
	return false
}

// Implement other user-related handlers as needed (update, delete, list, etc.)
