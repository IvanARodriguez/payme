package models

import "gorm.io/gorm"

func RunMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&Business{})
	return err
}
