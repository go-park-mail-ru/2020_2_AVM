package models

type (
	Profile struct {
		Id uint64 `json:"id"`
		Login string `json:"login"`
		Email string `json:"email"`
		Name string `json:"name"`
		Surname string `json:"surname"`
		Password string `json:"password"`
		Avatar   string `json:"avatar"`
		Cookie string `json:"-"`
	}
)