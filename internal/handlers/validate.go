package handlers

import (
	"go-and-htmx/internal/app"
	"regexp"

	"github.com/labstack/echo/v4"
)

func ValidateName(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)
	validName := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	formData := c.FormValue("name")

	if validName.MatchString(formData) {
		ctx.Validation.NameError = ""
		ctx.Validation.NameInput = formData
		return c.Render(200, "name_validation", ctx)
	} else {
		ctx.Validation.NameInput = formData
		ctx.Validation.NameError = "Invalid name. Only letters and spaces are allowed."
		return c.Render(200, "name_validation", ctx)
	}
}

func ValidateEmail(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)
	validEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	formData := c.FormValue("email")

	if validEmail.MatchString(formData) {
		ctx.Validation.EmailError = ""
		ctx.Validation.EmailInput = formData
		return c.Render(200, "email_validation", ctx)
	} else {
		ctx.Validation.EmailInput = formData
		ctx.Validation.EmailError = "Invalid email format."
		return c.Render(200, "email_validation", ctx)
	}
}
