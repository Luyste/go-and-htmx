package handlers

import (
	"github.com/labstack/echo/v4"
)

type Count struct {
	Counter int
}

func Welcome(c echo.Context) error {
	ctr := c.Get("counter").(*Count)
	return c.Render(200, "index", ctr)
}

func Increment(c echo.Context) error {
	// Get count, typecast into pointer struct (?) and increment
	ctr := c.Get("counter").(*Count)
	ctr.Counter++

	return c.Render(200, "counter", ctr)
}
