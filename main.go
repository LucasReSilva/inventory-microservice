package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db, _ := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	db.AutoMigrate(&Stock{})
}
