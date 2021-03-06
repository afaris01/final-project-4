package responses

import "time"

type UserRegisterResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUpdateResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}
