package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/go-park-mail-ru/2020_2_AVM/profile"
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ProfileHandler struct {
	useCase profile.ProfileUsecase
}

func NewProfileHandler (uC profile.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		useCase: uC,
	}
}

func (h *ProfileHandler) Signup(c echo.Context) (err error) {
	prof := new(models.Profile)
	if err = c.Bind(prof); err != nil {
		return
	}

	err = h.useCase.CreateProfile(prof)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, prof)

}

func (h *ProfileHandler) Signin(c echo.Context) (err error) {
	prof := new(models.Profile)
	expiration := time.Now().Add(8 * time.Hour)
	if err = c.Bind(prof); err != nil{
		return c.JSON(http.StatusBadRequest, prof)
	}

	baseProfile, err  := h.useCase.GetProfile(&prof.Login)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if prof.Password == baseProfile.Password {
		id := shortuuid.New()
		cookie := http.Cookie{
			Name:    "session_id",
			Value: id,
			Expires: expiration,
			HttpOnly: true,
		}
		c.SetCookie(&cookie)
		h.useCase.SetCookieToProfile(baseProfile, &cookie)
	} else {
		return c.JSON(http.StatusBadRequest, "Wrong password")
	}

	return c.JSON(http.StatusOK, prof)
}

func (h *ProfileHandler) Logout(c echo.Context) (err error) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	prof, err := h.useCase.GetProfileWithCookie(cookie)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	emptyCookie := new(http.Cookie)
	h.useCase.SetCookieToProfile(prof, emptyCookie)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "okk")
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

func (h *ProfileHandler) ProfileEditAvatar(c echo.Context) (err error) {
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, err)
	}
	prof, err :=  h.useCase.GetProfileWithCookie(cookie)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		userIdInt := int(prof.Id)

		err, filename := h.UploadAvatar(file, userIdInt)
		if err != nil {
			fmt.Println(err)
		} else {
			h.useCase.ProfileAvatarUpdate(prof, &filename)
		}
	}

	return c.JSON(http.StatusOK, "OK")
}


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
