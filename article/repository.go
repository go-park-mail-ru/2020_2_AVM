package article

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
)

type ArticleRepository interface {
	CreateArticle( article *models.Article ) error
	DeleteArticle( article *models.Article ) error
	GetArticlesByName( title *string ) ( []*models.Article, error )
	GetArticlesByAuthorId( authorId uint64 ) ( []*models.Article, error )
	GetArticlesByCategory( category *string ) ( []*models.Article, error )
	GetAllCategories() ( []*models.Category, error )
	CreateCategory( category models.Category ) error
	GetArticlesByTag( tag *string ) ( []*models.Article, error )
	CreateTag( tag models.Tag ) error
}

