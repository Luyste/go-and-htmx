package handlers

import (
	"go-and-htmx/internal/app"
	"strings"

	"github.com/labstack/echo/v4"
)

// if its a direct request (normal): render path/page
// if its a hx request: render path/index (only page content)

func Route(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	// Properly check HTMX header — Header.Get returns a string, compare to "true".
	if strings.EqualFold(c.Request().Header.Get("HX-Request"), "true") {
		// handle HX (partial) request
	}

	c.Path()

	// this is the base layout compiled with all components and route specific index template
	return c.Render(200, "index", ctx)
}
