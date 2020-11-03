package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_AVM/article"
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/go-park-mail-ru/2020_2_AVM/profile"
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type ProfileHandler struct {
	useCaseArt article.ArticleUsecase
	useCaseProf profile.ProfileUsecase
}

func NewProfileHandler (uCA article.ArticleUsecase, uCP profile.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		useCaseArt: uCA,
		useCaseProf: uCP,
	}
}

func (h *ProfileHandler) Signup(c echo.Context) (err error) {
	prof := new(models.Profile)
	if err = c.Bind(prof); err != nil {
		return
	}

	err = h.useCaseProf.CreateProfile(prof)

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

	baseProfile, err  := h.useCaseProf.GetProfile(&prof.Login)
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
		cookie_string := cookie.Value
		h.useCaseProf.SetCookieToProfile(baseProfile, &cookie_string)
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
	cookie_string := cookie.Value
	prof, err := h.useCaseProf.GetProfileWithCookie(&cookie_string)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	cookie_empty := ""
	h.useCaseProf.SetCookieToProfile(prof, &cookie_empty)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "okk")
}

func (h *ProfileHandler) Profile(c echo.Context) (err error) {
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, "bad")
	}

	cookie_string := cookie.Value
	answer, err := h.useCaseProf.GetProfileWithCookie(&cookie_string)
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
	cookie_string := cookie.Value
	profile, err := h.useCaseProf.GetProfileWithCookie(&cookie_string)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	if err = c.Bind(newProfile); err != nil {
		return
	}

	err = h.useCaseProf.UpdateProfile(profile, newProfile)

	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, newProfile)
}

func (h *ProfileHandler) SubscribeProfileToCategory(c echo.Context) (err error) {
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, "bad")
	}
	cookie_string := cookie.Value
	profile, err := h.useCaseProf.GetProfileWithCookie(&cookie_string)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	category := new(models.Category)
	if err = c.Bind(category); err != nil {
		return
	}
	category.Id, err = h.useCaseArt.GetCategoryID(&category.CategoryTitle)

	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.useCaseProf.SubscribeToCategory(profile, category)

	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	return nil
}

func (h *ProfileHandler) ProfileEditAvatar(c echo.Context) (err error) {
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, err)
	}
	cookie_string := cookie.Value
	prof, err :=  h.useCaseProf.GetProfileWithCookie(&cookie_string)
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
			h.useCaseProf.ProfileAvatarUpdate(prof, &filename)
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

	name := shortuuid.New() + "image"
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
