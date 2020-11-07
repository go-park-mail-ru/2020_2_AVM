package http

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/go-park-mail-ru/2020_2_AVM/article"
	"github.com/go-park-mail-ru/2020_2_AVM/profile"
	"strconv"
	"github.com/labstack/echo"
	"net/http"
)

type ArticleHandler struct {
	useCaseArt article.ArticleUsecase
	useCaseProf profile.ProfileUsecase
}

func NewAricleHandler (uCA article.ArticleUsecase, uCP profile.ProfileUsecase) *ArticleHandler {
	return &ArticleHandler{
		useCaseArt: uCA,
		useCaseProf: uCP,
	}
}


func (h *ArticleHandler) CreateArticle(c echo.Context) (err error) {
	art := new(models.Article)

	cookie, err := c.Cookie("session_id")
	if err = c.Bind(art); err != nil || cookie == nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	cookie_string := cookie.Value
	if prof, err := h.useCaseProf.GetProfileWithCookie(&cookie_string); err !=nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		art.AuthorID = prof.Id
	}
	categoryName := c.QueryParam("category_name")

	if categoryID, err := h.useCaseArt.GetCategoryID(&categoryName); err !=nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		art.CategoryID = categoryID
	}
	tags := new(models.Tags)
	if err = c.Bind(tags); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.useCaseArt.CreateArticle(art)
	h.useCaseArt.
	//Тэги
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, art)
}

func (h *ArticleHandler) ArticleByAuthor(c echo.Context) (err error) {
	key := c.Param("author")
	id, _ := strconv.Atoi(key)
	articles, err := h.useCaseArt.GetArticlesByAuthorId(uint64(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) SubscribedArticles(c echo.Context) (err error) {
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, "bad")
	}
	cookie_string := cookie.Value
	profile, err := h.useCaseProf.GetProfileWithCookie(&cookie_string)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := h.useCaseArt.GetArticlesBySubscribe(profile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}
func (h *ArticleHandler) ArticlesByTag(c echo.Context) (err error) {

	tag := new(models.Tag)

	if err = c.Bind(tag); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	articles, err := h.useCaseArt.GetArticlesByTag(&tag.TagTitle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) ArticlesByCategory(c echo.Context) (err error) {

	category := new(models.Category)
	if err = c.Bind(category); err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}
	articles, err := h.useCaseArt.GetArticlesByCategory(&category.CategoryTitle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, articles)
}


func (h *ArticleHandler) SubscribeToCategory(c echo.Context) (err error) {

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

	if err = c.Bind(category); err != nil || cookie == nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	category.Id, err = h.useCaseArt.GetCategoryID(&category.CategoryTitle)

	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}
	h.useCaseProf.SubscribeToCategory(profile, category)

	return nil
}