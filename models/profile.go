package models


type (
	Profile struct {
		Id string `json:"id"`
		Login string `json:"login"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
)
