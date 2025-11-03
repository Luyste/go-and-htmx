package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type Template struct {
	Name string
}

// BuildTemplateName creates a middleware that build the template name for the Render funciton.
func BuildTemplateName(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		logger := c.Logger()
		requestPath := strings.TrimPrefix(c.Request().URL.Path, "/")
		template := ""

		logger.Debugf("Request path: %v", requestPath)

		if strings.EqualFold(c.Request().Header.Get("HX-Request"), "true") {
			// handle HX (partial) request
			logger.Debug("HX request")
			if requestPath == "" {
				template = "index"
			} else {
				template = requestPath
			}
		} else {
			logger.Debug("Normal request")
			if requestPath == "" {
				template = "page"
			} else {
				template = requestPath
			}
		}

		logger.Debugf("template: %v", template)
		c.Set("template", template)

		return next(c)
	}
}
