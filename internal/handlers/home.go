package handlers

import (
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	ctx := c.Get("data").(*Context)
	return c.Render(200, "index", ctx)
}
