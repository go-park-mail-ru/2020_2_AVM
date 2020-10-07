package handlers

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) CreateArticle(c echo.Context) (err error) {
	art := new(models.Article)

	cookie, err := c.Cookie("session_id")
	if err = c.Bind(art); err != nil {
		return
	}
	art.Author = h.logInIds[cookie.Value]
	if err == http.ErrNoCookie {
		return c.JSON(http.StatusBadRequest, "bad")
	}

	art.Id = string(h.GetNewArcticleId())

	h.Articles = append(h.Articles, *art)
	return c.JSON(http.StatusCreated, art)
}

func (h *Handler) ArticleByAuthor(c echo.Context) (err error) {

	key := c.Param("author")
	articles := []models.Article{}
	for _, article := range h.Articles {
		if article.Author == key {
			articles = append(articles, article)
		}
	}

	return c.JSON(http.StatusOK, articles)
}

