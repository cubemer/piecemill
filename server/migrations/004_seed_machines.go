package migrations

import (
	"github.com/pocketbase/pocketbase/core"
)

func init() {
	core.AppMigrations.Register(func(txApp core.App) error {
		collection, err := txApp.FindCollectionByNameOrId("machines")
		if err != nil {
			return err
		}

		// Only seed if no machines exist yet (idempotent).
		count, err := txApp.CountRecords("machines")
		if err != nil {
			return err
		}
		if count > 0 {
			return nil
		}

		for i := 1; i <= 12; i++ {
			record := core.NewRecord(collection)
			record.Set("machine_id", i)
			record.Set("speed", 25)
			if err := txApp.Save(record); err != nil {
				return err
			}
		}

		return nil
	}, func(txApp core.App) error {
		// Down: delete all seeded machine records.
		records, err := txApp.FindAllRecords("machines")
		if err != nil {
			return err
		}
		for _, r := range records {
			if err := txApp.Delete(r); err != nil {
				return err
			}
		}
		return nil
	})
}
