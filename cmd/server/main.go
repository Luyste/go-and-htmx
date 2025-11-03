package main

import (
	"go-and-htmx/internal/app"
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
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)

	ctx := app.Context{Counter: 0}

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
	e.GET("/blog", handlers.Blog)

	// api
	e.POST("/increment", handlers.Increment)

	// start server
	e.Logger.Fatal(e.Start("localhost:42069"))
}
