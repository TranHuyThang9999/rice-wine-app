package adapters

import (
	"log"
	"rice-wine-shop/core/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPgsql() *gorm.DB {
	db, err := gorm.Open(postgres.Open(configs.Get().DataSource), &gorm.Config{})
	if err != nil {
		log.Fatal(err, "Failed to connect to the database")
		return nil
	}
	return db
}
