package profile

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"net/http"
)

type ProfileUsecase interface {
	CreateProfile( profile *models.Profile ) error
	DeleteProfile( profile *models.Profile ) error
	GetProfile( login *string ) ( *models.Profile, error )
	UpdateProfile( profile *models.Profile, profileNew *models.Profile) error
	ProfileAvatarUpdate ( profile *models.Profile, avatarPath *string) error
	SetCookieToProfile (profile *models.Profile, cookie *http.Cookie) error
	GetProfileWithCookie(cookie *http.Cookie) ( *models.Profile, error )
}