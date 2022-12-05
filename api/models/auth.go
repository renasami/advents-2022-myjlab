package models

type AuthRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type Auth struct {
	Id       string
	Username string
	Email    string
	Password string
}
