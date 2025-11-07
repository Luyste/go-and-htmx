package handlers

import (
	"go-and-htmx/internal/app"
	"strings"

	"github.com/labstack/echo/v4"
)

func Fragment(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)
	fragmentName := strings.Trim(c.Param("name"), ":")

	return c.Render(200, fragmentName, ctx)
}
