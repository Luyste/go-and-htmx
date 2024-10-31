package handlers

import (
	"github.com/labstack/echo/v4"
)

type Count struct {
	Counter int
}

type ContactInfo struct {
	Name  string
	Email string
}

type ContactList = []ContactInfo

type Context struct {
	Count    Count
	Contacts ContactList
}

func NewContact(name, email string) ContactInfo {
	return ContactInfo{
		Name:  name,
		Email: email,
	}
}

func Welcome(c echo.Context) error {
	ctx := c.Get("data").(*Context)
	return c.Render(200, "index", ctx)
}

func Increment(c echo.Context) error {
	// Get count, typecast into pointer struct (?) and increment
	ctx := c.Get("data").(*Context)
	ctx.Count.Counter++
	return c.Render(200, "counter", ctx)
}

func SaveContact(c echo.Context) error {
	ctx := c.Get("data").(*Context)

	name := c.FormValue("name")
	email := c.FormValue("email")
	nc := NewContact(name, email)

	ctx.Contacts = append(ctx.Contacts, nc)

	return c.Render(200, "form-display", ctx)
}
