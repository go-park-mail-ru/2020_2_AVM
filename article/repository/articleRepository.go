package repository

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"

)

type ArticleRepository struct {
	Atricles []models.Article
	articleId int
}

func (r *ArticleRepository) GetNewArcticleId() (int) {
	r.articleId += 1
	return r.articleId
}

func NewHandler() (*ArticleRepository) {
	return &ArticleRepository{nil, 0}
}


type ArcticleNotFound struct{}

func (t ArcticleNotFound) Error() string {
	return "Articles not found!"
}
type UnuniqueArticle struct{}

func (t UnuniqueArticle) Error() string {
	return "Ununique Article Data!"
}



func (r *ArticleRepository) CreateArticle( article *models.Article ) error {
	for _, art := range r.Atricles {
		if art.AuthorID == article.AuthorID && art.Title == article.Title {
			return UnuniqueArticle{}
		}
	}
	article.Id = uint64(r.GetNewArcticleId())
	r.Atricles = append(r.Atricles, *article)

	return nil
}
func (r *ArticleRepository) DeleteArticle( article *models.Article ) error {
	for i, art := range r.Atricles {
		if art.Id == article.Id {
			r.Atricles = append(r.Atricles[:i], r.Atricles[i + 1:]...)
			return nil
		}
	}

	return ArcticleNotFound{}
}

func (r *ArticleRepository) GetArticlesByName( title *string ) ( []*models.Article, error ) {
	var result = make([]*models.Article, 0)

	for _, art := range r.Atricles {
		if art.Title == *title {
			result = append(result, &art)
		}
	}
	if len(result) == 0 {
		return nil, ArcticleNotFound{}
	}

	return result, nil
}

func (r *ArticleRepository) GetArticlesByAuthorId( authorId uint64 ) ( []*models.Article, error ) {
	var result = make([]*models.Article, 0)

	for _, art := range r.Atricles {
		if art.AuthorID == authorId {
			result = append(result, &art)
		}
	}
	if len(result) == 0 {
		return nil, ArcticleNotFound{}
	}

	return result, nil
}



