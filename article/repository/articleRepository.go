package repository

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"gorm.io/gorm"
	"sync"
)

type ArticleRepository struct {
	conn *gorm.DB
	mute *sync.RWMutex
}

func NewAricleRepository(db *gorm.DB, mt *sync.RWMutex) *ArticleRepository {
	return &ArticleRepository{conn: db,
		mute: mt}
}

type ArcticleNotFound struct{}

func (t ArcticleNotFound) Error() string {
	return "Articles not found!"
}

type UnuniqueArticle struct{}

func (t UnuniqueArticle) Error() string {
	return "Ununique Article Data!"
}

type CategoryNotFound struct{}

func (t CategoryNotFound) Error() string {
	return "Category not found!"
}

func (adb *ArticleRepository) CreateArticle(article *models.Article) error {
	var err error
	adb.mute.Lock()
	{
		err = adb.conn.Table("articles").Create(article).Error
	}
	adb.mute.Unlock()

	return err
}
func (adb *ArticleRepository) DeleteArticle(article *models.Article) error {
	var err error
	adb.mute.Lock()
	{
		err = adb.conn.Table("articles").Delete(article).Error
	}
	adb.mute.Unlock()

	return err
}

func (adb *ArticleRepository) GetArticlesByName(title *string) ([]*models.Article, error) {
	var result []*models.Article

	var err error
	//var rows *sql.Rows
	adb.mute.RLock()
	{
		err = adb.conn.Table("articles").Where("article_title = ?", title).Find(result).Error
		/*	rows, err = adb.conn.Table("article").Where("title = ?", title).Rows()
			defer rows.Close()

			for rows.Next() {
				article := new(models.Article)
				adb.conn.ScanRows(rows, article)
				result = append(result, article)
			}*/
	}
	adb.mute.RUnlock()

	return result, err
}

func (adb *ArticleRepository) GetArticlesByAuthorId(authorId uint64) ([]*models.Article, error) {
	var result []*models.Article

	var err error
	//var rows *sql.Rows
	adb.mute.RLock()
	{
		err = adb.conn.Table("articles").Where("authorid = ?", authorId).Find(result).Error
		/*	rows, err = adb.conn.Table("article").Where("authorid = ?", authorId).Rows()
			defer rows.Close()

			for rows.Next() {
				article := new(models.Article)
				adb.conn.ScanRows(rows, article)
				result = append(result, article)
			}*/
	}
	adb.mute.RUnlock()

	return result, err
}

func (adb *ArticleRepository) GetArticlesByCategory(category *string) ([]*models.Article, error) {
	var result []*models.Article
	//var rows *sql.Rows
	var err error
	adb.mute.RLock()
	{
		var ctgr = new(models.Category)
		err = adb.conn.Table("category").Where("title = ?", category).First(ctgr).Error
		if err == nil {
			err = adb.conn.Table("articles").Where("categoryid = ?", ctgr.Id).Find(result).Error
			/*	rows, err = adb.conn.Table("article").Where("authorid = ?", authorId).Rows()
				defer rows.Close()

				for rows.Next() {
					article := new(models.Article)
					adb.conn.ScanRows(rows, article)
					result = append(result, article)
				}*/
		}
	}
	adb.mute.RUnlock()

	return result, err
}

func (adb *ArticleRepository) GetAllCategories() ([]*models.Category, error) {

	return nil, nil
}

func (adb *ArticleRepository) CreateCategory(category models.Category) error {
	var err error
	adb.mute.Lock()
	{
		err = adb.conn.Table("category").Create(category).Error
	}
	adb.mute.Unlock()

	return err
}

func (adb *ArticleRepository) GetArticlesById(id uint64) (*models.Article, error) {
	result := new(models.Article)

	var err error

	adb.mute.RLock()
	{
		err = adb.conn.Table("articles").Where("id = ?", id).First(result).Error
	}
	adb.mute.RUnlock()

	return result, err
}

func (adb *ArticleRepository) GetArticlesByTag(tag *string) ([]*models.Article, error) {
	var result []*models.Article
	var tagArticles []*models.TagArticle
	//var rows *sql.Rows
	var err error
	adb.mute.RLock()
	{
		var tg = new(models.Tag)
		err = adb.conn.Table("tag").Where("title = ?", tag).First(tg).Error //получаем id тэга
		if err == nil {
			err = adb.conn.Table("tag_article").Where("tagid = ?", tg.Id).Find(tagArticles).Error
			if err == nil {
				for _, tagArt := range tagArticles {
					article, err := adb.GetArticlesById(tagArt.ArticleID)
					if err == nil {
						result = append(result, article)
					}
				}
			}
			/*	rows, err = adb.conn.Table("article").Where("authorid = ?", authorId).Rows()
				defer rows.Close()

				for rows.Next() {
					article := new(models.Article)
					adb.conn.ScanRows(rows, article)
					result = append(result, article)
				}*/

		}
	}
	adb.mute.RUnlock()

	return result, err

}

func (adb *ArticleRepository) CreateTag(tag models.Tag) error {
	var err error
	adb.mute.Lock()
	{
		err = adb.conn.Table("tag").Create(tag).Error
	}
	adb.mute.Unlock()

	return err
}

func (adb *ArticleRepository) GetCategoryID(title *string) (uint64, error) {
	category := new(models.Category)

	var err error
	adb.mute.RLock()
	{
		err = adb.conn.Table("category").Where("category_title = ?", title).First(category).Error
	}
	adb.mute.RUnlock()

	return category.Id, err
}

func (adb *ArticleRepository) GetTagID(title *string) (uint64, error) {
	tag := new(models.Tag)

	var err error
	adb.mute.RLock()
	{
		err = adb.conn.Table("tag").Where("tag_title = ?", title).First(tag).Error
	}
	adb.mute.RUnlock()

	return tag.Id, err
}

func (adb *ArticleRepository) GetArticleIdByNameAndAuthorId (title *string, authorid uint64) (uint64, error) {
	article := new(models.Article)

	var err error
	adb.mute.RLock()
	{
		err = adb.conn.Table("article").Where("article_title = ? AND authorid = ?", title, authorid).First(article).Error
	}
	adb.mute.RUnlock()

	return article.Id, err
}

func (adb *ArticleRepository) GetSubscribedCategories(profile *models.Profile) ([]*models.Category, error) {
	var result []*models.Category

	var err error
	adb.mute.RLock()
	{
		err = adb.conn.Table("category_follow").Where("profileid = ?", profile.Id).Find(result).Error
	}
	adb.mute.RUnlock()

	return result, err
}

func (adb *ArticleRepository) LinkTagAndArticle(tagid uint64, articleid uint64) error {
	tagarticle := new(models.TagArticle)
	tagarticle.ArticleID = articleid
	tagarticle.TagID = tagid
	var err error
	adb.mute.Lock()
	{
		err = adb.conn.Table("tag_article").Create(tagarticle).Error
	}
	adb.mute.Unlock()

	return err
}
