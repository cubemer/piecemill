package main

import (
	"embed"
	"io/fs"
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "piecemill/migrations"
)

//go:embed all:dist
var distFS embed.FS

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Register custom /recommend endpoint before the static handler.
		se.Router.GET("/recommend", func(e *core.RequestEvent) error {
			return handleRecommend(e)
		})

		// Serve embedded frontend files.
		distContent, err := fs.Sub(distFS, "dist")
		if err != nil {
			return err
		}

		se.Router.GET("/{path...}", apis.Static(distContent, true))

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
