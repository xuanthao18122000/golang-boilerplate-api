package models

import "gorm.io/gorm"

type Posts struct {
	Id     uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name   *string `json:"name"`
	Status *int    `json:"status"`
}

func MigratePosts(db *gorm.DB) error {
	err := db.AutoMigrate(&Posts{})
	return err
}
