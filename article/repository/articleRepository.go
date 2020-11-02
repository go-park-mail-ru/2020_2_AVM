package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"gorm.io/gorm"
	"sync"
)

type ArticleRepository struct {
	conn   *gorm.DB
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

func (adb *ArticleRepository) CreateArticle( article *models.Article ) error {
	var err error
	adb.mute.Lock()
	{
		err = adb.conn.Table("article").Create(article).Error
	}
	adb.mute.Unlock()

	return err
}
func (adb *ArticleRepository) DeleteArticle( article *models.Article ) error {
	var err error
	adb.mute.Lock()
	{
		err = adb.conn.Table("article").Create(article).Error
	}
	adb.mute.Unlock()

	return err
}

func (adb *ArticleRepository) GetArticlesByName( title *string ) ( []*models.Article, error ) {
	var result []*models.Article

	var err error
	//var rows *sql.Rows
	adb.mute.RLock()
	{
		err =  adb.conn.Table("article").Where("title = ?", title).Create(result).Error
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

func (adb *ArticleRepository) GetArticlesByAuthorId( authorId uint64 ) ( []*models.Article, error ) {
	var result []*models.Article

	var err error
	//var rows *sql.Rows
	adb.mute.RLock()
	{
		err =  adb.conn.Table("article").Where("authorid = ?", authorId).Create(result).Error
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

func (adb *ArticleRepository) GetArticlesByCategory( category *string ) ( []*models.Article, error ) {
	var result []*models.Article
	//var rows *sql.Rows
	var err error
	adb.mute.RLock()
	{
		var ctgr = new(models.Category)
		err = adb.conn.Table("category").Where("title = ?", category).First(ctgr).Error
		if err == nil {
			err = adb.conn.Table("article").Where("categoryid = ?", ctgr.Id).Create(result).Error
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

func (adb *ArticleRepository) GetAllCategories() ( []*models.Category, error ) {

}

func (adb *ArticleRepository) CreateCategory( category models.Category ) error {
	var err error
	adb.mute.Lock()
	{
		err = adb.conn.Table("category").Create(category).Error
	}
	adb.mute.Unlock()

	return err
}

func (adb *ArticleRepository) GetArticlesByTag( tag *string ) ( []*models.Article, error ) {
	var result []*models.Article
	//var rows *sql.Rows
	var err error
	adb.mute.RLock()
	{
		var tg = new(models.Tag)
		err = adb.conn.Table("tag").Where("title = ?", tag).First(tg).Error
		if err == nil {
			err = adb.conn.Table("tag").Where("categoryid = ?", tg.Id).Create(result).Error
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

func (adb *ArticleRepository) CreateTag( tag models.Tag ) error {
	var err error
	adb.mute.Lock()
	{
		err = adb.conn.Table("tag").Create(tag).Error
	}
	adb.mute.Unlock()

	return err
}

