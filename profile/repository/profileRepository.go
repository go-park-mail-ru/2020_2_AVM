package repository

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"gorm.io/gorm"
	"net/http"
)

type ProfileRepository struct {
	conn   *gorm.DB
}
func NewProfileDatabase(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{conn: db}
}

type ProfileNotFound struct{}

func (t ProfileNotFound) Error() string {
	return "Profile not found!"
}

type UnuniqueProfileData struct{}

func (t UnuniqueProfileData) Error() string {
	return "Profile already exists!"
}


func (udb *ProfileRepository) CreateProfile(profile *models.Profile) (err error) {
	return udb.conn.Table("profile_repositories").Create(profile).Error
}
func (udb *ProfileRepository) DeleteProfile( profile *models.Profile ) error {
	return udb.conn.Table("profile_repositories").Delete(profile).Error
}

func (udb *ProfileRepository) GetProfile( login *string ) ( *models.Profile, error ) {
	profile := new(models.Profile)
	err := udb.conn.Table("profile_repositories").Where("login = ?", login).First(profile).Error

	return profile, err
}

func (udb *ProfileRepository) UpdateProfile( profile *models.Profile, profileNew *models.Profile) error {
	prof := new(models.Profile)
	err := udb.conn.Table("profile_repositories").Where("id = ?", profile.Id).First(prof).Error
	if err != nil {
		return err
	}
	if profileNew.Login != "" {
		prof.Login = profileNew.Login
	}
	if profileNew.Email != "" {
		prof.Email= profileNew.Email
	}
	if profileNew.Password != "" {
		prof.Password = profileNew.Password
	}

	if profileNew.Avatar != "" {
		prof.Avatar = profileNew.Avatar
	}

	if profileNew.Name != "" {
		prof.Name = profileNew.Name
	}

	if profileNew.Surname != "" {
		prof.Surname = profileNew.Surname
	}
	return udb.conn.Table("profile_repositories").Save(prof).Error

}

func (udb *ProfileRepository) GetProfileWithCookie(cookie *http.Cookie) ( *models.Profile, error ) {
	profile := new(models.Profile)
	err := udb.conn.Table("profile_repositories").Where("Cookie = ?", cookie).First(profile).Error

	return profile, err
}

func (udb *ProfileRepository) SetCookieToProfile (profile *models.Profile, cookie *http.Cookie) error {
	prof := new(models.Profile)
	err := udb.conn.Table("profile_repositories").Where("Id = ?", profile.Id).First(prof).Error
	if err != nil {
		return err
	}
	prof.Cookie = *cookie

	return udb.conn.Table("profile_repositories").Save(prof).Error
}