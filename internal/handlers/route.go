package handlers

import (
	"go-and-htmx/internal/app"

	"github.com/labstack/echo/v4"
)

// if its a direct request (normal): render path/page
// if its a hx request: render path/index (only page content)

func Route(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)
	template := c.Get("template").(string)

	// this is the base layout compiled with all components and route specific index template
	return c.Render(200, template, ctx)
}
