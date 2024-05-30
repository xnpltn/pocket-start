package application

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func NewApp() *pocketbase.PocketBase {
	app := pocketbase.New()
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		LoadRoutes(e, app)
		return nil
	})

	return app
}
