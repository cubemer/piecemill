package migrations

import (
	"github.com/pocketbase/pocketbase/core"
)

func init() {
	core.AppMigrations.Register(func(txApp core.App) error {
		collection := core.NewBaseCollection("machines")

		collection.Fields.Add(&core.NumberField{
			Name:     "machine_id",
			Required: true,
			Min:      ptrFloat(1),
			Max:      ptrFloat(12),
		})

		collection.Fields.Add(&core.NumberField{
			Name:     "speed",
			Required: true,
			Min:      ptrFloat(18),
			Max:      ptrFloat(25),
		})

		return txApp.Save(collection)
	}, func(txApp core.App) error {
		collection, err := txApp.FindCollectionByNameOrId("machines")
		if err != nil {
			return err
		}
		return txApp.Delete(collection)
	})
}
