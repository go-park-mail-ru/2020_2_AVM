package http

import (
	"github.com/go-park-mail-ru/2020_2_AVM/article"
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/go-park-mail-ru/2020_2_AVM/profile"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	useCaseArt  article.ArticleUsecase
	useCaseProf profile.ProfileUsecase
}

func NewAricleHandler(uCA article.ArticleUsecase, uCP profile.ProfileUsecase) *ArticleHandler {
	return &ArticleHandler{
		useCaseArt:  uCA,
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
	prof, err := h.useCaseProf.GetProfileWithCookie(&cookie_string)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		art.AuthorID = prof.Id
	}

	category := new(models.Category)
	category.CategoryTitle = c.FormValue("category_title")

	if categoryID, err := h.useCaseArt.GetCategoryID(&category.CategoryTitle); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		art.CategoryID = categoryID
	}

	err = h.useCaseArt.CreateArticle(art)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	//нужно получить id статьи методом ниже, чтобы потом связывать его через Link
	//GetArticleIdByNameAndAuthorId
	articleid, err := h.useCaseArt.GetArticleIdByNameAndAuthorId(&art.ArticleTitle, prof.Id)

	tags := new(models.Tags) // нужно получить массив тэгов и потом для каждого вызвать Link
	tags.TagsValues = c.FormValue("tags")
	for _, tag := range tags.TagsValues {
		tagid, err  := h.useCaseArt.GetTagID(&tag.TagTitle)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		h.useCaseArt.LinkTagAndArticle(tagid, articleid)
	}

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
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := h.useCaseArt.GetArticlesBySubscribe(profile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}
func (h *ArticleHandler) ArticlesByTag(c echo.Context) (err error) {
	tagname := c.Param("tag")

	articles, err := h.useCaseArt.GetArticlesByTag(&tagname)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) ArticlesByCategory(c echo.Context) (err error) {

	category := new(models.Category)
	if err = c.Bind(category); err != nil {
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
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	category := new(models.Category)

	if err = c.Bind(category); err != nil || cookie == nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	category.Id, err = h.useCaseArt.GetCategoryID(&category.CategoryTitle)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	h.useCaseProf.SubscribeToCategory(profile, category)

	return nil
}
