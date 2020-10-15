package http

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/go-park-mail-ru/2020_2_AVM/prodile"
	"github.com/go-park-mail-ru/2020_2_AVM/article"

	"github.com/labstack/echo"
	"net/http"
)

type ArticleHandler struct {
	useCase article.ArticleUsecase
}

func NewAricleHandler (useCase article.ArticleUsecase) *ArticleHandler {
	return &ArticleHandler{
		useCase: useCase,
	}
}


func (h *ArticleHandler) CreateArticle(c echo.Context) (err error) {
	art := new(models.Article)

	cookie, err := c.Cookie("session_id")
	if err = c.Bind(art); err != nil || cookie == nil {
		return
	}

	if id, ok := h.useCase.GetProfileWithCookie(cookie); err == http.ErrNoCookie || !ok {
		return c.JSON(http.StatusBadRequest, "Unlogined user")
	} else {
		art.AuthorID = id
	}
	err = h.useCase.CreateAticle(art)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, art)
}

func (h *ArticleHandler) ArticleByAuthor(c echo.Context) (err error) {
	key := c.Param("author")
	articles, err := h.useCase.GetArticlesByAuthorId(key)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, articles)
}