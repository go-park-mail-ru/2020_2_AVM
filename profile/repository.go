package profile

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"net/http"
)

type ProfileRepository interface {
	CreateProfile( profile *models.Profile ) error
	DeleteProfile( profile *models.Profile ) error
	GetProfile( login *string ) ( *models.Profile, error )
	UpdateProfile( profile *models.Profile, profileNew *models.Profile ) error
	GetProfileWithCookie(cookie *http.Cookie) ( *models.Profile, error )
}