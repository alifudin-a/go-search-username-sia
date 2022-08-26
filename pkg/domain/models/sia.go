package models

// User sia struct
type UserSIA struct {
	ID       string `json:"id,omitempty" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password,omitempty" db:"password"`
	Email    string `json:"email" db:"email"`
	FullName string `json:"full_name" db:"full_name"`
}
