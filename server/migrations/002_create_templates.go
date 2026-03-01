package migrations

import (
	"github.com/pocketbase/pocketbase/core"
)

func init() {
	core.AppMigrations.Register(func(txApp core.App) error {
		collection := core.NewBaseCollection("templates")

		collection.Fields.Add(&core.TextField{
			Name:     "template_id",
			Required: true,
		})

		collection.Fields.Add(&core.NumberField{
			Name:     "cut_time_at_25",
			Required: true,
			Min:      ptrFloat(0),
		})

		return txApp.Save(collection)
	}, func(txApp core.App) error {
		collection, err := txApp.FindCollectionByNameOrId("templates")
		if err != nil {
			return err
		}
		return txApp.Delete(collection)
	})
}
