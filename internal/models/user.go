package models

// User represents a user of the car rental system.
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"` // Note: You should securely hash passwords in a real system.
    // Add more user-related fields as needed
}

// NewUser creates a new User instance.
func NewUser(id int, username, email, password string) *User {
    return &User{
        ID:       id,
        Username: username,
        Email:    email,
        Password: password,
    }
}
