package config

import (
	"assignment-3/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase(dbHost string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbHost), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Order{}, &models.Item{}, &models.User{})

	return db, nil
}
