package handlers

import (
	"go-and-htmx/internal/app"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)
	return c.Render(200, "base", ctx)
}
