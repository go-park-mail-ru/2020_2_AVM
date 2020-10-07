package handlers

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	prof := new(models.Profile)
	if err = c.Bind(prof); err != nil {
		return
	}
	prof.Id = strconv.Itoa(h.GetNewUserId())
	h.Profiles = append(h.Profiles, *prof)
	return c.JSON(http.StatusCreated, prof)

}

func (h *Handler) ProfileEdit(c echo.Context) (err error) {
	new_profile := new(models.Profile)
	cookie, err := c.Cookie("session_id")
	user_id := h.logInIds[cookie.Value]

	if err = c.Bind(&new_profile); err != nil {
		return
	}
	for i, profile := range h.Profiles {
		if profile.Id == user_id {
			h.Profiles[i].ConfirmChanges(*new_profile)
		}
	}

	return c.JSON(http.StatusOK, new_profile)
}
