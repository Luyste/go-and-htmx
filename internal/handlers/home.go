package handlers

import (
	"go-and-htmx/internal/app"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	return c.Render(200, "index", ctx)
}

func Back(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	c.Logger().Debugf("request context: %v", ctx)

	return c.Render(200, "todo", ctx)
}
