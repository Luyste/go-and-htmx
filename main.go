package main

import (
	u "go-and-htmx/internal/handlers"
	"go-and-htmx/tools"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// initiate echo
	e := echo.New()

	// initialize middleware
	e.Use(middleware.Logger())

	// seed to see some data
	contacts := u.ContactList{
		u.NewContact("Jop", "Jop@gmail.com"),
		u.NewContact("Kevin", "Kevin@gmail.com"),
	}

	pageData := u.Contacts{
		Contacts: contacts,
	}
	formData := u.FormData{}

	ctx := u.Context{
		// Data to render the page
		PageData: pageData,

		// Data to render the form
		FormData: formData,
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("data", &ctx)
			return next(c)
		}
	})

	e.Renderer = render.NewTemplate()

	// routers
	e.GET("/", u.Home)
	e.POST("/save-contact", u.SaveContact)

	// start server
	e.Logger.Fatal(e.Start("localhost:42069"))
}
