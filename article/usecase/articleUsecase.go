package usecase

import (
	"github.com/go-park-mail-ru/2020_2_AVM/article"
	"github.com/go-park-mail-ru/2020_2_AVM/models"
)

type ArticleUseCase struct {
	DBConnArt article.ArticleRepository
}

func NewArticleUseCase( dbConnArt article.ArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{
		DBConnArt: dbConnArt,
	}
}

func (h *ArticleUseCase) CreateArticle( article *models.Article ) error {
	return h.DBConnArt.CreateArticle(article)

}
func (h *ArticleUseCase) DeleteArticle( article *models.Article ) error {
	return h.DBConnArt.DeleteArticle(article)
}
func (h *ArticleUseCase) GetArticlesByName( title *string ) ( []*models.Article, error ) {
	return h.DBConnArt.GetArticlesByName(title)
}
func (h *ArticleUseCase) GetArticlesByAuthorId( authorId uint64 ) ( []*models.Article, error ) {
	return h.DBConnArt.GetArticlesByAuthorId(authorId)
}

func (h *ArticleUseCase) GetArticlesByCategory( category *string ) ( []*models.Article, error ) {
	return h.DBConnArt.GetArticlesByCategory(category)
}

func (h *ArticleUseCase) GetArticlesByTag( tag *string ) ( []*models.Article, error ) {
	return h.DBConnArt.GetArticlesByTag(tag)
}


