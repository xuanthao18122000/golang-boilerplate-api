package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	Id     uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name   *string `json:"name"`
	Status *int    `json:"status"`
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
