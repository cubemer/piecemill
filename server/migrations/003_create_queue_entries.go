package migrations

import (
	"github.com/pocketbase/pocketbase/core"
)

func init() {
	core.AppMigrations.Register(func(txApp core.App) error {
		collection := core.NewBaseCollection("queue_entries")

		collection.Fields.Add(&core.TextField{
			Name:     "template_id",
			Required: true,
		})

		collection.Fields.Add(&core.NumberField{
			Name:     "machine_id",
			Required: true,
			Min:      ptrFloat(1),
			Max:      ptrFloat(12),
		})

		collection.Fields.Add(&core.NumberField{
			Name:     "quantity",
			Required: true,
			Min:      ptrFloat(1),
		})

		return txApp.Save(collection)
	}, func(txApp core.App) error {
		collection, err := txApp.FindCollectionByNameOrId("queue_entries")
		if err != nil {
			return err
		}
		return txApp.Delete(collection)
	})
}
