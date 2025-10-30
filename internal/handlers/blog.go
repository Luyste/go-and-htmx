package handlers

import (
	"go-and-htmx/internal/app"

	"github.com/labstack/echo/v4"
)

func Blog(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	if c.Request().Header.Get("HX-Request") == "true" {
		return c.Render(200, "content", ctx)
	}

	return c.Render(200, "base", ctx)
}
