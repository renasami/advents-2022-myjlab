package models

type AuthRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}


type Auth struct {
	Username string 
	Email string 
	Password string 
}

type UserDb struct {
	Id string 
	Username string 
	Email string 
	Password string 
}