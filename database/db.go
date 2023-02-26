package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDb() {
	conn := "host=localhost user=root password=root dbname=root port=5437 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conn))
	if err != nil {
		log.Panic("Errode conexão ", err)
	}
}
