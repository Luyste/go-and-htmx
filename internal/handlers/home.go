package handlers

import (
	"go-and-htmx/internal/app"
	"go-and-htmx/internal/utils"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	if utils.IsHTMX(c) {
		return c.Render(200, "content_home", ctx)
	}

	return c.Render(200, "layout_home", ctx)
}
