package main

import (
	articleDelivery "github.com/go-park-mail-ru/2020_2_AVM/article/delivery/http"
	articleRepository "github.com/go-park-mail-ru/2020_2_AVM/article/repository"
	articleUseCase "github.com/go-park-mail-ru/2020_2_AVM/article/usecase"
	profileDelivery "github.com/go-park-mail-ru/2020_2_AVM/profile/delivery/http"
	profileRepository "github.com/go-park-mail-ru/2020_2_AVM/profile/repository"
	profileUseCase "github.com/go-park-mail-ru/2020_2_AVM/profile/usecase"
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
	artRepository := articleRepository.NewAricleRepository()
	profRepository := profileRepository.NewProfileRepository()

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
		AllowOrigins: []string{"http://localhost:1323", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, CSRFHeader},
	}))


	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())


	// Routes
	e.POST("/article", serverConfig.ArticleHandler.CreateArticle)

	e.GET("/article/author/:author", serverConfig.ArticleHandler.ArticleByAuthor)
	e.GET("/avatar", serverConfig.profileHandler.AvatarDefault)
	e.GET("/avatar/title/:name", serverConfig.profileHandler.Avatar)
	e.GET("/profile", serverConfig.profileHandler.Profile)
	e.PUT("/setting/avatar", serverConfig.profileHandler.ProfileEditAvatar)
	e.PUT("/setting", serverConfig.profileHandler.ProfileEdit)
	e.POST("/signup", serverConfig.profileHandler.Signup)
	e.POST("/signin", serverConfig.profileHandler.Signin)
	e.POST("/logout", serverConfig.profileHandler.Logout)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}