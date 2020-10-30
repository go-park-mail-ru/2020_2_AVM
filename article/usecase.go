package article

import "github.com/go-park-mail-ru/2020_2_AVM/models"

type ArticleUsecase interface {
	CreateArticle( article *models.Article ) error
	DeleteArticle( article *models.Article ) error
	GetArticlesByName( title *string ) ( []*models.Article, error )
	GetArticlesByAuthorId( authorId uint64 ) ( []*models.Article, error )
}

