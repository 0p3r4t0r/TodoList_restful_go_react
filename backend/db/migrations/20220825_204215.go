package main

import (
	"backend/db"
	"backend/db/models"
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		fmt.Print("Could not connect to DB!")
	}

	m := gormigrate.New(database, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20220825_204215",
			Migrate: func(tx *gorm.DB) error {

				return tx.AutoMigrate(&models.Task{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("tasks")
			},
		},
	})

	if err = m.Migrate(); err != nil {
		fmt.Print("Migration failed!")
	}

	fmt.Print("Migration successful!")
}
