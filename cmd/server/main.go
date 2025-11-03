package main

import (
	"go-and-htmx/internal/app"
	"go-and-htmx/internal/handlers"
	"go-and-htmx/internal/render"

	customMiddleware "go-and-htmx/internal/middleware"

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

	templates, err := render.LoadTemplates()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Renderer = templates

	// routes
	home := e.Group("/")
	blog := e.Group("/blog")
	form := e.Group("/form")

	home.Use(customMiddleware.BuildTemplateName)
	blog.Use(customMiddleware.BuildTemplateName)
	form.Use(customMiddleware.BuildTemplateName)

	homeFragments := home.Group("/f")
	blogFragments := blog.Group("/f")
	formFragments := home.Group("/f")

	home.GET("", handlers.Route)
	homeFragments.GET("", handlers.Fragment)

	blog.GET("", handlers.Route)
	blogFragments.GET("", handlers.Fragment)

	form.GET("", handlers.Route)
	formFragments.GET("", handlers.Fragment)

	//
	e.POST("/increment", handlers.Increment)

	// start server
	e.Logger.Fatal(e.Start("localhost:42069"))
}
