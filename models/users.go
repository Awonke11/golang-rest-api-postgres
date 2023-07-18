package models

import (
	"gorm.io/gorm"
)

type Users struct {
	user_id int    `json:"user_id`
	name    string `json:"name"`
	age     int    `json:"age"`
	phone   string `json:"phone"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
