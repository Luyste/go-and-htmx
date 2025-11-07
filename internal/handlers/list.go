package handlers

import (
	"go-and-htmx/internal/app"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AddToList(c echo.Context) error {
	item := c.FormValue("item")
	ctx := c.Get("data").(*app.Context)

	itemId := "item-" + strings.ReplaceAll(uuid.NewString(), "-", "")
	ctx.List = append(ctx.List, app.Item{Id: itemId, Item: item})

	c.Logger().Infof("String len: $v", len(ctx.List))

	if len(ctx.List) > 0 {
		return c.Render(200, "todo_list_reset_input_field_share_button", ctx)
	} else {
		return c.Render(200, "todo_list", ctx)
	}

}

func RemoveFromList(c echo.Context) error {
	ctx := c.Get("data").(*app.Context)
	id := strings.Trim(c.Param("id"), ":")

	newList := []app.Item{}

	for _, listItem := range ctx.List {
		if listItem.Id != id {
			newList = append(newList, listItem)
		}
	}

	ctx.List = newList

	c.Logger().Debugf("List len: %v", len(ctx.List))

	if len(ctx.List) == 0 {
		return c.Render(200, "todo_list_reset_input_field", ctx)
	} else {
		return c.Render(200, "todo_list_reset_input_field_share_button", ctx)
	}

}
