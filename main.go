package main

import (
	"github.com/go-park-mail-ru/2020_2_AVM/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	CSRFHeader := "X-CSRF-TOKEN"

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1323", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, CSRFHeader},
	}))


	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	// Initialize handler
	h := handlers.NewHandler()
	// Routes
	e.POST("/article", h.CreateArticle)

	e.GET("/article/:author", h.ArticleByAuthor)
	e.POST("/signup", h.Signup)
	e.POST("/signin", h.Login)
	e.POST("/logout", h.Logout)
	e.POST("/setting", h.ProfileEdit)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}