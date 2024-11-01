package main

import (
	"go-and-htmx/internal/app"
	"go-and-htmx/internal/handlers"
	"go-and-htmx/tools"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	// my applications state is captured in the app.Context struct. This holds all relevant information
	ctx := app.Context{
		FormData: app.NewFormData(),
		DisplayData: app.DisplayData{
			app.NewContact("Jop", "jop@email.com"),
			app.NewContact("kevin", "kevin@email.com"),
		},
		Counter: 0,
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("data", &ctx)
			return next(c)
		}
	})

	e.Renderer = render.NewTemplate()

	e.GET("/", handlers.Index)
	e.POST("/add", handlers.Add)
	e.POST("/increment", handlers.Increment)

	e.Logger.Fatal(e.Start("localhost:42069"))
}
