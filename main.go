package main

import (
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

	// Set counter as 0 and let middleware add it to context
	ctr := handlers.Count{Counter: 0}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Note the pointer on ctr, allowing the handler to modify the original struct
			c.Set("counter", &ctr)
			return next(c)
		}
	})

	e.Renderer = render.NewTemplate()

	// routers
	e.GET("/", handlers.Welcome)
	e.POST("/increment", handlers.Increment)

	// start server
	e.Logger.Fatal(e.Start(":42069"))
}
