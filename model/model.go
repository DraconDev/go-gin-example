package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// func GetUser(id int) (*User, error) {
// 	// Logic to retrieve user from database or other storage
// }

// ... other model functions
