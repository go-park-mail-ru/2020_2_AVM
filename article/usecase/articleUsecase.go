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
func (h *ArticleUseCase) GetCategoryID (title *string) (uint64, error){
	return h.DBConnArt.GetCategoryID(title)
}
func (h *ArticleUseCase) GetTagID (title *string) (uint64, error){
	return h.DBConnArt.GetTagID(title)
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
func (h *ArticleUseCase) GetArticlesByTags( tag models.Tags ) ( []*models.Article, error ) {
	var result []*models.Article
	for _, tag := range tag.TagsValues {
		buff, err :=  h.DBConnArt.GetArticlesByTag(&tag.TagTitle)
		if err != nil {
			return nil, err
		}
		result = append(result, buff...)
	}
	return result, nil
}

func (h *ArticleUseCase) GetArticlesBySubscribe( profile *models.Profile ) ( []*models.Article, error ) {

	return nil, nil
}



