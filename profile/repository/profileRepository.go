package repository

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"net/http"
)

type ProfileRepository struct {
	Profiles []models.Profile
	userId int
}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{
		Profiles: []models.Profile{},
		userId: 0,
	}
}


func (r *ProfileRepository) GetNewUserId() (int) {
	r.userId += 1
	return r.userId
}


func NewHandler() (*ProfileRepository) {
	return &ProfileRepository{nil,  0}
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
			return UnuniqueProfileData{}
		}
	}
	profile.Id = uint64(r.GetNewUserId())
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

func (r *ProfileRepository) UpdateProfile( profile *models.Profile, profileNew *models.Profile) error {
	if _, err := r.GetProfile(&profile.Login); err != nil {
		return ProfileNotFound{}
	}

	for i, prof := range r.Profiles {
		if prof.Id == profile.Id {
			if profileNew.Login != "" {
				r.Profiles[i].Login = profileNew.Login
			}
			if profileNew.Email != "" {
				r.Profiles[i].Email = profileNew.Email
			}
			if profileNew.Password != "" {
				r.Profiles[i].Password = profileNew.Password
			}

			if profileNew.Avatar != "" {
				r.Profiles[i].Avatar = profileNew.Avatar
			}

			if profileNew.Name != "" {
				r.Profiles[i].Name = profileNew.Name
			}

			if profileNew.Surname != "" {
				r.Profiles[i].Surname = profileNew.Surname
			}
			return nil
		}
	}


	return nil
}

func (r *ProfileRepository) GetProfileWithCookie(cookie *http.Cookie) ( *models.Profile, error ) {
	for _, prof := range r.Profiles {
		if prof.Cookie.Value == cookie.Value {
			return &prof, nil

		}
	}

	return nil, ProfileNotFound{}
}

func (r *ProfileRepository) SetCookieToProfile (profile *models.Profile, cookie *http.Cookie) error {
	for i, prof := range r.Profiles {
		if prof.Id == profile.Id {
			r.Profiles[i].Cookie = * cookie
			return nil
		}
	}
	return ProfileNotFound{}
}