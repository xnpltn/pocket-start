package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/xnpltn/pocket-start/templates"
)

func RenderView(c echo.Context, view templ.Component, layoutPath string) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return view.Render(c.Request().Context(), c.Response().Writer)
	}
	return templates.Layout(layoutPath).Render(c.Request().Context(), c.Response().Writer)
}
