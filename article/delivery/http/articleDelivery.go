package http

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"github.com/go-park-mail-ru/2020_2_AVM/profile"
	"github.com/go-park-mail-ru/2020_2_AVM/article"
	"strconv"

	"github.com/labstack/echo"
	"net/http"
)

type ArticleHandler struct {
	useCase article.ArticleUsecase
	profileRepository profile.ProfileRepository
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

	if prof, err := h.profileRepository.GetProfileWithCookie(cookie); err !=nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		art.AuthorID = prof.Id
	}
	err = h.useCase.CreateArticle(art)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, art)
}

func (h *ArticleHandler) ArticleByAuthor(c echo.Context) (err error) {
	key := c.Param("author")
	id, _ := strconv.Atoi(key)
	articles, err := h.useCase.GetArticlesByAuthorId(uint64(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, articles)
}