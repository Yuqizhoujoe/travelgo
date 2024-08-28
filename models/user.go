package models

type AddUser struct {
	Email string `json:"email"`
}

type AddUserResponse struct {
	Success bool `json:"success"`
}
