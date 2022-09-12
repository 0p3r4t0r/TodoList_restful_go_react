package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
 * Connection setup for running within docker.
 * Set host=localhost and port=5500 to connect from your machine as in ConnectLocal below.
 */
func Connect() (*gorm.DB, error) {

	dsn := "host=db user=test_user password=test dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

/*
 * Connection setup for running migrations from local machine.
 * Set host=db and port=5432 to connect from within docker as in Connect above.
 */
func ConnectLocal() (*gorm.DB, error) {

	dsn := "host=localhost user=test_user password=test dbname=test_db port=5500 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
