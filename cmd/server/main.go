package main

import (
	"go-and-htmx/internal/app"
	context "go-and-htmx/internal/app"
	"go-and-htmx/internal/handlers"
	render "go-and-htmx/tools"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// initiate echo
	e := echo.New()

	// initialize middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}; uri=${uri}; status=${status}; error=${error}\n",
	}))
	e.Logger.SetLevel(log.DEBUG)

	ctx := app.Context{List: []context.Item{}}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("data", &ctx)
			return next(c)
		}
	})

	// render stylesheets
	e.Static("/static", "web/static")

	e.Renderer = render.NewTemplate()

	// fragments
	e.GET("/fragments:name", handlers.Fragment)

	// routes
	e.GET("/", handlers.Home)
	e.GET("/back", handlers.Back)

	// api
	list := e.Group("/list")
	list.POST("/add", handlers.AddToList)
	list.DELETE("/remove:id", handlers.RemoveFromList)

	share := e.Group("/share")
	share.GET("/view", handlers.ShareView)
	share.POST("/submit", handlers.ShareSubmit)

	validate := e.Group("/validate")
	validate.POST("/name", handlers.ValidateName)
	validate.POST("/email", handlers.ValidateEmail)

	// start server
	e.Logger.Fatal(e.Start("localhost:42069"))
}
