package models

type (
	Profile struct {
		Id string `json:"id"`
		Login string `json:"login"`
		Email string `json:"email"`
		Password string `json:"password"`
		Avatar   string `json:"avatar"`
	}
)
func (p Profile) ConfirmChanges(profile Profile) ()  {
	if profile.Login != "" {
		p.Login = profile.Login
	}
	if profile.Email != "" {
		p.Email = profile.Email
	}
	if profile.Password != "" {
		p.Password = profile.Password
	}
}
