package handlers

import (
	"go-and-htmx/internal/app"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	// this is the base layout compiled with all components and route specific index template
	return c.Render(200, "index", ctx)
}
