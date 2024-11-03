package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-and-htmx/internal/app"
)

func Add(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)

	name := c.FormValue("name")
	email := c.FormValue("email")

	// now lets build in some validation and error codes.
	if ctx.HasEmail(email) {
		//If email exist, we want to show an error message that email already exists
		// For this, we're going to render the form input fields with provided input
		// + an error message.

		// Lets use the FormData struct
		// Note that we don't want to store this error in the state of the app, therefore
		// create a copy of context.

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

		fmt.Printf(">>>> old email detected <<<<")
		fmt.Printf(">>>> ectx %+v\n", fd)
		return c.Render(200, "Form", fd)
	}
	// it looks like the 500 error is because of a data race.
	// if email is not yet present, add it.
	ctx.DisplayData.Contacts = append(ctx.DisplayData.Contacts, app.NewContact(name, email))
	return c.Render(200, "Display", ctx.DisplayData)
}
