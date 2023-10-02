package models

import "gorm.io/gorm"

type Posts struct {
	id		uint `gorm:"primary key;autoIncrement" json:"id"`
	name	*string `json:"name"`
	status  *int `json:"status"`
}

func MigratePosts(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}