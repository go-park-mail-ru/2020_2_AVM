package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/go-park-mail-ru/2020_2_AVM/profile"
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type ProfileHandler struct {
	useCase profile.ProfileUsecase
}

func NewProfileHandler (useCase profile.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		useCase: useCase,
	}
}

func (h *ProfileHandler) Signup(c echo.Context) (err error) {
	prof := new(models.Profile)
	if err = c.Bind(prof); err != nil {
		return
	}

	err = h.useCase.CreateProfile(prof)

	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, prof)

}

func (h *ProfileHandler) Profile(c echo.Context) (err error) {
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, "bad")
	}

	answer, err := h.useCase.GetProfileWithCookie(cookie)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, answer)
}

func (h *ProfileHandler) ProfileEdit(c echo.Context) (err error) {
	newProfile := new(models.Profile)
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, "bad")
	}
	profile, err := h.useCase.GetProfileWithCookie(cookie)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	if err = c.Bind(newProfile); err != nil {
		return
	}

	err = h.useCase.UpdateProfile(profile, newProfile)

	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, newProfile)
}
/*
func (h *ProfileHandler) ProfileEditAvatar(c echo.Context) (err error) {
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, "bad")
	}
	prof, err :=  h.useCase.GetProfileWithCookie(cookie)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad")
	} else {
		userIdInt := int(prof.Id)

		err, filename := h.UploadAvatar(file, userIdInt)
		if err != nil {
			fmt.Println(err)
		} else {
			for i, profile := range h.us {
				if profile.Id == userIdInt {
					h.Profiles[i].Avatar = filename
				}
			}
		}
	}

	return c.JSON(http.StatusOK, "OK")
}*/


func (h *ProfileHandler) AvatarDefault(c echo.Context) (err error) { // rework
	return c.File("./static/avatars/default_avatar.png")
}

func (h *ProfileHandler) Avatar(c echo.Context) (err error) { // rework
	filename := c.Param("name")
	return c.File("./static/avatars/" + filename)
}

func (h *ProfileHandler) UploadAvatar(file *multipart.FileHeader, userID int) (err error, filename string) {
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
