package handlers

import (
	"go-and-htmx/internal/app"

	"github.com/labstack/echo/v4"
)

func ShareView(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	ctx.Validation.NameError = ""
	ctx.Validation.NameInput = ""
	ctx.Validation.EmailError = ""
	ctx.Validation.EmailInput = ""

	return c.Render(200, "share", ctx)
}

func ShareSubmit(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	c.Logger().Debugf("request context: %+v", ctx)

	return c.Render(200, "share", ctx)
}
