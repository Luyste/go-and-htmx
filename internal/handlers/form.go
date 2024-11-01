package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func NewContact(name, email string) ContactInfo {
	return ContactInfo{
		Name:  name,
		Email: email,
	}
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func (c *Context) hasEmail(email string) bool {
	for _, contact := range c.PageData.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func SaveContact(c echo.Context) error {
	ctx := c.Get("data").(*Context)

	name := c.FormValue("name")
	email := c.FormValue("email")

	if ctx.hasEmail(email) {
		fd := NewFormData()

		fd.Values["name"] = name
		fd.Values["email"] = email
		fd.Errors["email"] = "Email already exists!"

		ectx := *ctx
		ectx.FormData = fd

		fmt.Printf("FormData: ", ectx)

		return c.Render(200, "index", ectx)
	}

	nc := NewContact(name, email)
	ctx.PageData.Contacts = append(ctx.PageData.Contacts, nc)

	return c.Render(200, "index", ctx)
}
