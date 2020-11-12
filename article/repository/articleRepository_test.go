package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_AVM/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type dbMock struct {
	db      *gorm.DB
	mock    sqlmock.Sqlmock
}

var (
	testArticle = models.Article{
		Id:           1,
		ArticleTitle: "ArticleTitle",
		Description:  "Description",
		Content:      "Content",
		CategoryID:   1,
		AuthorID:     1,
	}

	testCategory = models.Category{
		Id:            1,
		CategoryTitle: "CategoryTitle",
	}
)

// https://github.com/go-gorm/gorm/issues/3565

func TestCreateArticle(t *testing.T) {
	s := &dbMock{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	dsn := "host=localhost user=avm_user password=qwerty123 dbname=avmvc port=5432 sslmode=disable"
	dialector := postgres.New(postgres.Config{
		DSN:                  dsn,
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}

	defer db.Close()

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "articles" ("article_title","description","content","categoryid","authorid")
					VALUES ($1,$2,$3,$4,$5) RETURNING "articles"."id"`)).
		WithArgs(
			testArticle.ArticleTitle,
			testArticle.Description,
			testArticle.Content,
			testArticle.CategoryID,
			testArticle.AuthorID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(testArticle.Id))

	s.mock.ExpectCommit()

	//if err = s.db.Create(testArticle).Error; err != nil {
	//	t.Errorf("Failed to insert to gorm db, got error: %v", err)
	//}
	//
	//err = s.mock.ExpectationsWereMet()
	//if err != nil {
	//	t.Errorf("Failed to meet expectations, got error: %v", err)
	//}
}

func TestDeleteArticle(t *testing.T) {

}

func TestGetArticlesByName(t *testing.T) {

}

func TestGetArticlesByAuthorId(t *testing.T) {

}

func TestGetArticlesByCategory(t *testing.T) {

}

func TestGetAllCategories(t *testing.T) {

}

func TestCreateCategory(t *testing.T) {

}

func TestGetArticlesById(t *testing.T) {

}

func TestGetArticlesByTag(t *testing.T) {

}

func TestCreateTag(t *testing.T) {

}

func TestGetCategoryID(t *testing.T) {

}

func TestGetTagID(t *testing.T) {

}

func TestGetArticleIdByNameAndAuthorId(t *testing.T) {

}

func TestGetSubscribedCategories(t *testing.T) {

}

func TestLinkTagAndArticle(t *testing.T) {

}
