package handlers

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
	file, err := c.FormFile("avatar")
	if err != nil {
		fmt.Println(err)
	} else {
		user_id_int, _ := strconv.Atoi(user_id)
		err, _ := h.uploadAvatar(file, user_id_int)
		if err != nil {
			fmt.Println(err)
		}
	}



	return c.JSON(http.StatusOK, new_profile)
}

func (h *Handler) Avatar(c echo.Context) (err error) {
	filename := c.Param("name")

	if filename == "default_avatar.png" {
		return c.File("./default/default_avatar.png")
	}
	return c.File("./avatars/" + filename)
}

func (h *Handler) uploadAvatar(file *multipart.FileHeader, userID int) (err error, filename string) {
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	defer src.Close()

	hash := sha256.New()
	formattedTime := strings.Join(strings.Split(time.Now().String(), " "), "")
	formattedID := strconv.FormatUint(uint64(userID), 10)

	name := fmt.Sprintf("%x", hash.Sum([]byte(formattedTime+formattedID)))
	filename = name + ".jpeg"
	dst, err := os.Create("./avatars/" + filename)

	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println(err)
		return err, ""
	}

	h.Profiles[userID].Avatar = filename
	return nil, filename
}