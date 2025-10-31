package handlers

import (
	"go-and-htmx/internal/app"

	"github.com/labstack/echo/v4"
)

func Form(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	return c.Render(200, "form/index", ctx)
}
