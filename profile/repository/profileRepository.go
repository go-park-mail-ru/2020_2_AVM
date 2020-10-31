package repository

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"gorm.io/gorm"
	"sync"
)

type ProfileRepository struct {
	conn   *gorm.DB
	mute *sync.RWMutex
}

func NewProfileRepository(db *gorm.DB, mt *sync.RWMutex) *ProfileRepository {
	return &ProfileRepository{conn: db,
								mute: mt}
}

type ProfileNotFound struct{}

func (t ProfileNotFound) Error() string {
	return "Profile not found!"
}

type UnuniqueProfileData struct{}

func (t UnuniqueProfileData) Error() string {
	return "Profile already exists!"
}

func (udb *ProfileRepository) CreateProfile(profile *models.Profile) error {
	var err error
	udb.mute.Lock()
	{
		err = udb.conn.Table("user_profile").Create(profile).Error
	}
	udb.mute.Unlock()

	return err
}
func (udb *ProfileRepository) DeleteProfile( profile *models.Profile ) error {
	var err error
	udb.mute.Lock()
	{
		err =  udb.conn.Table("user_profile").Delete(profile).Error
	}
	udb.mute.Unlock()

	return err
}

func (udb *ProfileRepository) GetProfile( login *string ) ( *models.Profile, error ) {
	profile := new(models.Profile)

	var err error
	udb.mute.RLock()
	{
		err = udb.conn.Table("user_profile").Where("login = ?", login).First(profile).Error
	}
	udb.mute.RUnlock()

	return profile, err
}

func (udb *ProfileRepository) UpdateProfile( profile *models.Profile, profileNew *models.Profile) error {
	prof := new(models.Profile)
	var err error
	udb.mute.RLock()
	{
		err = udb.conn.Table("user_profile").Where("id = ?", profile.Id).First(prof).Error
	}
	udb.mute.RUnlock()

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

	udb.mute.Lock()
	{
		err = udb.conn.Table("user_profile").Save(prof).Error
	}
	udb.mute.Unlock()

	return err


}

func (udb *ProfileRepository) GetProfileWithCookie(cookie *string) ( *models.Profile, error ) {
	profile := new(models.Profile)

	var err error
	udb.mute.RLock()
	{
		err = udb.conn.Table("user_profile").Where("Cookie = ?", cookie).First(profile).Error
	}
	udb.mute.RUnlock()

	return profile, err
}

func (udb *ProfileRepository) SetCookieToProfile (profile *models.Profile, cookie *string) error {
	prof := new(models.Profile)
	var err error
	udb.mute.RLock()
	{
		err = udb.conn.Table("user_profile").Where("Id = ?", profile.Id).First(prof).Error
	}
	udb.mute.RUnlock()

	if err != nil {
		return err
	}
	prof.Cookie = *cookie

	udb.mute.Lock()
	{
		err = udb.conn.Table("user_profile").Save(prof).Error
	}
	udb.mute.Unlock()

	return err

}