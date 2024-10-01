package database

import (
	"GoSimpleAPI/src/app/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(CFG *config.Configuration) error {
	var err error

	DB, err = gorm.Open(postgres.Open(CFG.DatabaseUrl), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	return nil
}
