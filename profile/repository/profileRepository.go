package repository

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/labstack/echo"
	"mime/multipart"
	"net/http"
	"strconv"
	"fmt"
	"io"
	"os"
	"sync"
)

type ProfileRepository struct {
	Profiles []models.Profile
	logInIds map[string]string
}

type ProfileNotFound struct{}

func (t ProfileNotFound) Error() string {
	return "Profile not found!"
}

type UnuniqueProfileData struct{}

func (t UnuniqueProfileData) Error() string {
	return "Profile already exists!"
}

func (r *ProfileRepository) CreateProfile( profile *models.Profile ) error {

	for _, prof := range r.Profiles {
		if prof.Login == profile.Login || prof.Email == profile.Email {
			return nil
		}
			return UnuniqueProfileData{}
	}
	r.Profiles = append(r.Profiles, *profile)

	return nil
}

func (r *ProfileRepository) DeleteProfile( profile *models.Profile ) error {
	for i, prof := range r.Profiles {
		if prof.Id == profile.Id {
			r.Profiles = append(r.Profiles[:i], r.Profiles[i + 1:]...)
			return nil
		}
	}
	return ProfileNotFound{}
}

func (r *ProfileRepository) GetProfile( login *string ) ( *models.Profile, error ) {


	for _, prof := range r.Profiles {
		if prof.Login == *login {
			return &prof, nil
		}
	}
	return nil, ProfileNotFound{}
}

func (r *ProfileRepository) UpdateProfile( profile *models.Profile, name, surname, login, email, password, avatarPath string ) error {
	if _, err := r.GetProfile(&profile.Login); err != nil {
		return ProfileNotFound{}
	}

	if login != "" {
		profile.Login = login
	}
	if email != "" {
		profile.Email = email
	}
	if password != "" {
		profile.Password = password
	}

	if avatarPath != "" {
		profile.Avatar = avatarPath
	}

	if name != "" {
		profile.Name = name
	}

	if surname != "" {
		profile.Surname = surname
	}


}
func (r *ProfileRepository) GetProfileWithCookie(cookie *http.Cookie) ( *models.Profile, error ) {
	for _, prof := range r.Profiles {
		if prof.Id == r.logInIds[cookie.Value] {
			return &prof, nil
		}
	}

	return nil, ProfileNotFound{}
}