package handlers

import (
	"go-and-htmx/internal/app"

	"github.com/labstack/echo/v4"
)

func Increment(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)
	ctx.Counter.Count++

	return c.Render(200, "Counter", ctx.Counter)
}
