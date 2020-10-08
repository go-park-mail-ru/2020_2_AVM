package main

import (
	"github.com/go-park-mail-ru/2020_2_AVM/handlers"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	CSRFHeader := "X-CSRF-TOKEN"

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins: []string{"http://localhost:1323", "http://localhost:8082"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, CSRFHeader},
	}))


	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	// Initialize handler
	h := handlers.NewHandler()
	// Routes
	e.POST("/api/article", h.CreateArticle)

	e.GET("/api/article/:author", h.ArticleByAuthor)
	e.GET("/api/avatar/", h.AvatarDefault)
	e.GET("/api/avatar/:name", h.Avatar)
	e.GET("/api/profile", h.Profile)
	e.PUT("/api/setting/avatar", h.ProfileEditAvatar)
	e.PUT("/api/setting", h.ProfileEdit)
	e.POST("/api/signup", h.Signup)
	e.POST("/api/signin", h.Login)
	e.POST("/api/logout", h.Logout)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}