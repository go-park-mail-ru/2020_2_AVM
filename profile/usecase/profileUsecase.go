package usecase

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/go-park-mail-ru/2020_2_AVM/profile"
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

func (h *ProfileUseCase) GetProfileWithCookie(cookie *string)( *models.Profile, error){
	return h.DBConn.GetProfileWithCookie(cookie)
}


func (h *ProfileUseCase) UpdateProfile( profile *models.Profile, profileNew *models.Profile) error {
	return h.DBConn.UpdateProfile(profile, profileNew)
}
func (h *ProfileUseCase) ProfileAvatarUpdate ( profile *models.Profile, avatarPath *string) error {
	prof := new(models.Profile)
	prof.Avatar = *avatarPath
	return h.DBConn.UpdateProfile(profile, prof)
}

func (h *ProfileUseCase) SetCookieToProfile (profile *models.Profile, cookie *string) error {
	return h.DBConn.SetCookieToProfile(profile, cookie)
}
