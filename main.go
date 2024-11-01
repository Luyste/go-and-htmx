package main

import (
	"go-and-htmx/internal/app"
	"go-and-htmx/internal/handlers"
	"go-and-htmx/tools"

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

	e.Renderer = render.NewTemplate()

	// routers
	e.GET("/", handlers.Index)
	e.POST("/increment", handlers.Increment)

	// start server
	e.Logger.Fatal(e.Start("localhost:42069"))
}
