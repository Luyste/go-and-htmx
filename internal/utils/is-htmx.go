package utils

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func IsHTMX(c echo.Context) bool {
	return strings.EqualFold(c.Request().Header.Get("HX-Request"), "true")
}
