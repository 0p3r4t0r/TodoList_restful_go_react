package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	/*
	 * Connection setup for running within docker.
	 * Set host=localhost and port=5500 to connect from your machine.
	 */
	dsn := "host=db user=test_user password=test dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
