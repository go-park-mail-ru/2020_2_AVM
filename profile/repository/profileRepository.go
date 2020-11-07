package repository

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"gorm.io/gorm"
	"sync"
)

type ProfileRepository struct {
	conn *gorm.DB
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

func (pdb *ProfileRepository) CreateProfile(profile *models.Profile) error {
	var err error
	pdb.mute.Lock()
	{
		err = pdb.conn.Table("profiles").Create(profile).Error
	}
	pdb.mute.Unlock()

	return err
}
func (pdb *ProfileRepository) DeleteProfile(profile *models.Profile) error {
	var err error
	pdb.mute.Lock()
	{
		err = pdb.conn.Table("profiles").Delete(profile).Error
	}
	pdb.mute.Unlock()

	return err
}

func (pdb *ProfileRepository) GetProfile(login *string) (*models.Profile, error) {
	profile := new(models.Profile)

	var err error
	pdb.mute.RLock()
	{
		err = pdb.conn.Table("profiles").Where("login = ?", login).First(profile).Error
	}
	pdb.mute.RUnlock()

	return profile, err
}

func (pdb *ProfileRepository) UpdateProfile(profile *models.Profile, profileNew *models.Profile) error {
	prof := new(models.Profile)
	var err error
	pdb.mute.RLock()
	{
		err = pdb.conn.Table("profiles").Where("id = ?", profile.Id).First(prof).Error
	}
	pdb.mute.RUnlock()

	if err != nil {
		return err
	}
	if profileNew.Login != "" {
		prof.Login = profileNew.Login
	}
	if profileNew.Email != "" {
		prof.Email = profileNew.Email
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

	pdb.mute.Lock()
	{
		err = pdb.conn.Table("profiles").Save(prof).Error
	}
	pdb.mute.Unlock()

	return err

}

func (pdb *ProfileRepository) GetProfileWithCookie(cookie *string) (*models.Profile, error) {
	profile := new(models.Profile)

	var err error
	pdb.mute.RLock()
	{
		err = pdb.conn.Table("profiles").Where("Cookie = ?", cookie).First(profile).Error
	}
	pdb.mute.RUnlock()

	return profile, err
}

func (pdb *ProfileRepository) SetCookieToProfile(profile *models.Profile, cookie *string) error {
	prof := new(models.Profile)
	var err error
	pdb.mute.RLock()
	{
		err = pdb.conn.Table("profiles").Where("Id = ?", profile.Id).First(prof).Error
	}
	pdb.mute.RUnlock()

	if err != nil {
		return err
	}
	prof.Cookie = *cookie

	pdb.mute.Lock()
	{
		err = pdb.conn.Table("profiles").Save(prof).Error
	}
	pdb.mute.Unlock()

	return err

}

func (pdb *ProfileRepository) SubscribeToCategory(profile *models.Profile, category *models.Category) error {
	sub := new(models.CategoryFollow)
	sub.CategoryID = category.Id
	sub.ProfileID = profile.Id

	var err error
	pdb.mute.Lock()
	{
		err = pdb.conn.Table("category_follow").Create(sub).Error
	}
	pdb.mute.Unlock()

	return err
}
