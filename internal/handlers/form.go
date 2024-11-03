package handlers

import (
	"github.com/labstack/echo/v4"
	"go-and-htmx/internal/app"
)

func Add(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	name := c.FormValue("name")
	email := c.FormValue("email")

	if ctx.HasEmail(email) {
		// TODO: Set correct error code!

		fd := app.NewFormData()
		fd.Values["name"] = name
		fd.Values["email"] = email
		fd.Errors["email"] = "This email already exists!"

		// cheat our way, return 200 because 400 - 500 are not rendered by hmtx default
		// also we now only want to render the form, not the display

		// problem, we are telling the engine to only render "Form", but Form.html has an
		// target attribute pointing towards Display.
		// By default the requestee, will also get the response.

		// TODO: How do we render "Display" without hx-target="#display"????
		// TODO: How do we render "Form" without hx-target="this"????

		return c.Render(200, "Form", fd)
	}
	// it looks like the 500 error is because of a data race.
	// if email is not yet present, add it.
	ctx.DisplayData.Contacts = append(ctx.DisplayData.Contacts, app.NewContact(name, email))
	return c.Render(200, "Display", ctx.DisplayData)
}
