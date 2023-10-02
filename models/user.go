package models

import "gorm.io/gorm"

type Users struct {
	Id     uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name   *string `json:"name"`
	Status *int    `json:"status"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
