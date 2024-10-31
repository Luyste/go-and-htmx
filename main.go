package main

import (
	h "go-and-htmx/internal/handlers"
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
	count := h.Count{Counter: 0}
	contacts := h.ContactList{
		h.NewContact("Jop", "Jop@gmail.com"),
		h.NewContact("Kevin", "Kevin@gmail.com"),
	}
	ctx := h.Context{
		Count:    count,
		Contacts: contacts,
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Since we're now passing more data, we need a way to:
			// 1. Send everything in one go, but also be able to substract
			// necessary information from dataset.
			// create a map of key/value pairs which hold references to the
			// structs.

			c.Set("data", &ctx)
			return next(c)
		}
	})

	e.Renderer = render.NewTemplate()

	// routers
	e.GET("/", h.Welcome)
	e.POST("/increment", h.Increment)
	e.POST("/save-contact", h.SaveContact)

	// start server
	e.Logger.Fatal(e.Start(":42069"))
}
