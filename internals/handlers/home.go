package handlers

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/xnpltn/pocket-start/templates/views"
	"github.com/xnpltn/pocket-start/utils"
)

func HomeGet(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return utils.RenderView(c, views.HomeView("Home"), "/")
	}
}
