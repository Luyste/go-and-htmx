package handlers

import (
	"github.com/labstack/echo/v4"
	"go-and-htmx/internal/app"
)

func Index(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)
	return c.Render(200, "index", ctx)
}
