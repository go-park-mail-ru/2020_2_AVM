package handlers

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (h *Handler) CreateArticle(c echo.Context) (err error) {
	art := new(models.Article)

	cookie, err := c.Cookie("session_id")
	if err = c.Bind(art); err != nil || cookie == nil {
		return
	}

	if id, ok := h.logInIds[cookie.Value]; err == http.ErrNoCookie || !ok {
		return c.JSON(http.StatusBadRequest, "Unlogined user")
	} else {
		art.AuthorID = id
	}

	art.Id = strconv.Itoa(h.GetNewArcticleId())

	h.Articles = append(h.Articles, *art)
	return c.JSON(http.StatusCreated, art)
}

func (h *Handler) ArticleByAuthor(c echo.Context) (err error) {
	key := c.Param("author")
	articles := []models.Article{}

	for _, article := range h.Articles {
		if article.AuthorID == key {
			articles = append(articles, article)
		}
	}

	return c.JSON(http.StatusOK, articles)
}