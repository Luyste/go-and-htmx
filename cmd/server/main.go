package main

import (
	"go-and-htmx/internal/app"
	"go-and-htmx/internal/handlers"
	"go-and-htmx/internal/render"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// initiate echo
	e := echo.New()

	// initialize middleware
	e.Use(middleware.Logger())

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

	homeFragments := home.Group("/f")
	blogFragments := blog.Group("/f")
	formFragments := home.Group("/f")

	blog.Add("GET", "", handlers.Route)
	blogFragments.Add("GET", "", handlers.Fragment)

	home.Add("GET", "", handlers.Route)
	homeFragments.Add("GET", "", handlers.Fragment)

	form.Add("GET", "", handlers.Route)
	formFragments.Add("GET", "", handlers.Fragment)

	//
	e.POST("/increment", handlers.Increment)

	// start server
	e.Logger.Fatal(e.Start("localhost:42069"))
}
