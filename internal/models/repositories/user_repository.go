package repositories

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"github.com/Anubhav-Chauhan3367/CarBookingSystem.git/internal/models"
)

// UserRepositoryJSON is an implementation of UserRepository using JSON files.
type UserRepositoryJSON struct {
	filePath string
	mu       sync.Mutex // Mutex for concurrent access
}

func NewUserRepositoryJSON(filePath string) *UserRepositoryJSON {
	return &UserRepositoryJSON{filePath: filePath}
}

func (r *UserRepositoryJSON) GetUserByID(id int) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(r.filePath)
	fmt.Printf("data: %s\n", data)
	fmt.Printf("err: %s\n", err)
	if err != nil {
		return nil, err
	}

	var users []*models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user with ID %d not found", id)
}

func (r *UserRepositoryJSON) GetUserByUsername(username string) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var users []*models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user with username %s not found", username)
}

func (r *UserRepositoryJSON) CreateUser(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	var users []*models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return err
	}

	// Assign a unique ID to the new user
	if len(users) == 0 {
		user.ID = 1
	} else {
		lastUser := users[len(users)-1]
		user.ID = lastUser.ID + 1
	}

	users = append(users, user)

	newData, err := json.Marshal(users)
	if err != nil {
		return err
	}

	if err := os.WriteFile(r.filePath, newData, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// Implement UpdateUser and DeleteUser methods similarly
