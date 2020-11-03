package article

import "github.com/go-park-mail-ru/2020_2_AVM/models"

type ArticleUsecase interface {
	CreateArticle( article *models.Article ) error
	DeleteArticle( article *models.Article ) error
	GetCategoryID (title *string) (uint64, error)
	GetTagID (title *string) (uint64, error)
	GetArticlesByName( title *string ) ( []*models.Article, error )
	GetArticlesByAuthorId( authorId uint64 ) ( []*models.Article, error )
	GetArticlesByCategory( category *string ) ( []*models.Article, error )
	GetArticlesByTag( tag *string ) ( []*models.Article, error )
	GetArticlesBySubscribe( profile *models.Profile ) ( []*models.Article, error )
}

