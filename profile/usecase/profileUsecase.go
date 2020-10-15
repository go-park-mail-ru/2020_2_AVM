package usecase

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/go-park-mail-ru/2020_2_AVM/profile"
	"net/http"
)

type ProfileUseCase struct {
	DBConn profile.ProfileRepository
}


func NewProfileUseCase( dbConn profile.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{
		DBConn: dbConn,
	}
}


func (h *ProfileUseCase) CreateProfile( profile *models.Profile ) error {
	return h.DBConn.CreateProfile(profile)
}


func (h *ProfileUseCase) DeleteProfile( profile *models.Profile ) error {
	return h.DBConn.DeleteProfile(profile)
}


func (h *ProfileUseCase) GetProfile( login *string )( *models.Profile, error ) {
	return h.DBConn.GetProfile(login)
}

func (h *ProfileUseCase) GetProfileWithCookie(cookie *http.Cookie)( *models.Profile, error){
	return h.DBConn.GetProfileWithCookie(cookie)
}


func (h *ProfileUseCase) UpdateProfile( profile *models.Profile, profileNew *models.Profile) error {
	return h.DBConn.UpdateProfile(profile, profileNew)
}