package usecase

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
)

type ArticleUseCase struct {
	DBConn article.ArticleRepository
}

func NewArticleUseCase( dbConn article.ArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{
		DBConn: dbConn,
	}
}

func (h *ArticleUseCase) CreateArticle( article *models.Article ) error {
	return h.DBConn.CreateArticle(article)

}
func (h *ArticleUseCase) DeleteArticle( article *models.Article ) error {
	return h.DBConn.DeleteArticle(article)
}
func (h *ArticleUseCase) GetArticlesByName( title *string ) ( []*models.Article, error ) {
	return h.DBConn.GetArticlesByName(title)
}
func (h *ArticleUseCase) GetArticlesByAuthorId( authorId uint64 ) ( []*models.Article, error ) {
	return h.DBConn.GetArticlesByAuthorId(authorId)
}


