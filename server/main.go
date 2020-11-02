package main

import (
	model "github.com/go-park-mail-ru/2020_2_AVM/models"
	articleDelivery "github.com/go-park-mail-ru/2020_2_AVM/article/delivery/http"
	articleRepository "github.com/go-park-mail-ru/2020_2_AVM/article/repository"
	articleUseCase "github.com/go-park-mail-ru/2020_2_AVM/article/usecase"
	profileDelivery "github.com/go-park-mail-ru/2020_2_AVM/profile/delivery/http"
	profileRepository "github.com/go-park-mail-ru/2020_2_AVM/profile/repository"
	profileUseCase "github.com/go-park-mail-ru/2020_2_AVM/profile/usecase"
	"sync"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type ServerStruct struct{
	ArticleHandler *articleDelivery.ArticleHandler
	profileHandler *profileDelivery.ProfileHandler
	httpServer *http.Server
}

func configureAPI() *ServerStruct{
	mutex := sync.RWMutex{}

	dsn := "host=localhost user=avm_user password=qwerty123 dbname=avmvc port=5432 sslmode=disable"
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

	artHandler := articleDelivery.NewAricleHandler(artUseCase, profUseCase)
	profHandler := profileDelivery.NewProfileHandler(profUseCase)

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
		AllowOrigins: []string{"http://localhost:1323", "http://localhost:3000", "http://95.163.250.127:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, CSRFHeader},
	}))


	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())


	// Routes
	e.POST("/api/article", serverConfig.ArticleHandler.CreateArticle)

	e.GET("/api/article/author/:author", serverConfig.ArticleHandler.ArticleByAuthor)
	e.GET("/api/avatar", serverConfig.profileHandler.AvatarDefault)
	e.GET("/api/avatar/title/:name", serverConfig.profileHandler.Avatar)
	e.GET("/api/profile", serverConfig.profileHandler.Profile)
	e.PUT("/api/setting/avatar", serverConfig.profileHandler.ProfileEditAvatar)
	e.PUT("/api/setting", serverConfig.profileHandler.ProfileEdit)
	e.POST("/api/signup", serverConfig.profileHandler.Signup)
	e.POST("/api/signin", serverConfig.profileHandler.Signin)
	e.POST("/api/logout", serverConfig.profileHandler.Logout)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}