package handlers

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/labstack/echo"
	"mime/multipart"
	"net/http"
	"strconv"
	"fmt"
	"io"
	"os"

)

func (h *Handler) Signup(c echo.Context) (err error) {
	prof := new(models.Profile)
	if err = c.Bind(prof); err != nil {
		return
	}

	for _, profile := range h.Profiles {
		if profile.Login == prof.Login || profile.Email == profile.Email {
			return c.JSON(http.StatusBadRequest, "Ununique data")
		}
	}
	prof.Id = strconv.Itoa(h.GetNewUserId())
	h.Profiles = append(h.Profiles, *prof)
	return c.JSON(http.StatusCreated, prof)

}

func (h *Handler) Profile(c echo.Context) (err error) {
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, "bad")
	}
	userLogin, ok := h.logInIds[cookie.Value]
	if !ok {
		return c.JSON(http.StatusBadRequest, "Unlogged user")
	}
	answer := new(models.Profile)
	for _, profile := range h.Profiles {
		if profile.Login == userLogin {
			*answer = profile
		}
	}

	return c.JSON(http.StatusCreated, answer)
}

func (h *Handler) ProfileEdit(c echo.Context) (err error) {
	newProfile := new(models.Profile)
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, "bad")
	}
	userId, ok := h.logInIds[cookie.Value]
	if !ok {
		return c.JSON(http.StatusBadRequest, "Unlogged user")
	}

	if err = c.Bind(newProfile); err != nil {
		return
	}

	for i, profile := range h.Profiles {
		if profile.Id == userId {
			h.Profiles[i].ConfirmChanges(*newProfile)
		}
	}

	return c.JSON(http.StatusOK, newProfile)
}

func (h *Handler) ProfileEditAvatar(c echo.Context) (err error) {
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, "bad")
	}
	userId, ok :=  h.logInIds[cookie.Value]
	if !ok {
		return c.JSON(http.StatusBadRequest, "Unlogged user")
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad")
	} else {
		userIdInt, _ := strconv.Atoi(userId)

		err, filename := h.UploadAvatar(file, userIdInt)
		if err != nil {
			fmt.Println(err)
		} else {
			for i, profile := range h.Profiles {
				if profile.Id == userId {
					h.Profiles[i].Avatar = filename
				}
			}
		}
	}

	return c.JSON(http.StatusOK, "OK")
}


func (h *Handler) AvatarDefault(c echo.Context) (err error) { // rework
	return c.File("./static/avatars/default_avatar.png")
}

func (h *Handler) Avatar(c echo.Context) (err error) { // rework
	filename := c.Param("name")
	return c.File("./static/avatars/" + filename)
}

func (h *Handler) UploadAvatar(file *multipart.FileHeader, userID int) (err error, filename string) {
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	defer src.Close()

	name := strconv.Itoa(userID * 666) + "image"
	filename = name + ".jpeg"
	dst, err := os.Create("./static/avatars/" + filename)

	if err != nil {
		return err, ""
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err, ""
	}

	return nil, filename
}