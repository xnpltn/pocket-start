package application

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/xnpltn/pocket-start/internals/handlers"
)

func LoadRoutes(e *core.ServeEvent, app *pocketbase.PocketBase) {
	e.Router.Static("/static", "assets")
	loadViews(e, app)
}

func loadViews(e *core.ServeEvent, app *pocketbase.PocketBase) {
	e.Router.GET("/", handlers.HomeGet(app))
	e.Router.GET("/about", handlers.AboutGet(app))
}
