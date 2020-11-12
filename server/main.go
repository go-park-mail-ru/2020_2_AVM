package main

import (
	articleDelivery "github.com/go-park-mail-ru/2020_2_AVM/article/delivery/http"
	articleRepository "github.com/go-park-mail-ru/2020_2_AVM/article/repository"
	articleUseCase "github.com/go-park-mail-ru/2020_2_AVM/article/usecase"
	model "github.com/go-park-mail-ru/2020_2_AVM/models"
	profileDelivery "github.com/go-park-mail-ru/2020_2_AVM/profile/delivery/http"
	profileRepository "github.com/go-park-mail-ru/2020_2_AVM/profile/repository"
	profileUseCase "github.com/go-park-mail-ru/2020_2_AVM/profile/usecase"

	"github.com/microcosm-cc/bluemonday"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type ServerStruct struct {
	ArticleHandler *articleDelivery.ArticleHandler
	profileHandler *profileDelivery.ProfileHandler
	httpServer     *http.Server
}

func configureAPI() *ServerStruct {
	mutex := sync.RWMutex{}

	p := bluemonday.UGCPolicy()

	//dsn := "host=localhost user=avm_user password=qwerty123 dbname=avmvc port=5432 sslmode=disable"
	dsn := "host=localhost user=mark password=mark dbname=mark_avm_db port=5432 sslmode=disable"
	//dsn := "host=localhost user=postgres password=mark dbname=postgres_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&profileRepository.ProfileRepository{})

	db.Migrator().CreateTable(&model.Article{})
	db.Migrator().CreateTable(&model.Profile{})

	artRepository := articleRepository.NewAricleRepository(db, &mutex)
	profRepository := profileRepository.NewProfileRepository(db, &mutex)

	artUseCase := articleUseCase.NewArticleUseCase(artRepository)
	profUseCase := profileUseCase.NewProfileUseCase(profRepository)

	artHandler := articleDelivery.NewAricleHandler(artUseCase, profUseCase, p)
	profHandler := profileDelivery.NewProfileHandler(artUseCase, profUseCase, p)

	return &ServerStruct{
		ArticleHandler: artHandler,
		profileHandler: profHandler,
	}
}

func main() {
	e := echo.New()
	serverConfig := configureAPI()

	CSRFHeader := "X-CSRF-TOKEN"

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:1323", "http://localhost:8080", "http://95.163.250.127:8080"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, CSRFHeader},
	}))

	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())

	// Routes

	e.GET("/api/article/author/:author", serverConfig.ArticleHandler.ArticleByAuthor)
	e.GET("/api/article/tag/:tag", serverConfig.ArticleHandler.ArticlesByTag)
	e.GET("/api/article/category/:category", serverConfig.ArticleHandler.ArticlesByCategory)
	e.GET("/api/article/subscribe/", serverConfig.ArticleHandler.SubscribedArticles)
	e.GET("/api/avatar", serverConfig.profileHandler.AvatarDefault)
	e.GET("/api/avatar/title/:name", serverConfig.profileHandler.Avatar)
	e.GET("/api/profile", serverConfig.profileHandler.Profile)
	e.PUT("/api/setting/avatar", serverConfig.profileHandler.ProfileEditAvatar)
	e.PUT("/api/setting", serverConfig.profileHandler.ProfileEdit)
	e.POST("/api/article", serverConfig.ArticleHandler.CreateArticle)
	e.POST("/api/signup", serverConfig.profileHandler.Signup)
	e.POST("/api/signin", serverConfig.profileHandler.Signin)
	e.POST("/api/logout", serverConfig.profileHandler.Logout)
	e.POST("/api/subscribe", serverConfig.ArticleHandler.SubscribeToCategory)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
