package profile

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
)

type ProfileRepository interface {
	CreateProfile( profile *models.Profile ) error
	DeleteProfile( profile *models.Profile ) error
	GetProfile( login *string ) ( *models.Profile, error )
	UpdateProfile( profile *models.Profile, profileNew *models.Profile ) error
	SetCookieToProfile (profile *models.Profile, cookie *string) error
	GetProfileWithCookie(cookie *string) ( *models.Profile, error )
	SubscribeToCategory(profile *models.Profile, category *models.Category) error
}