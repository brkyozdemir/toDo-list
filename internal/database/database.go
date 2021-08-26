package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDB() *gorm.DB {
	dsn := os.Getenv("DB_CONNECTION") + os.Getenv("DB_NAME") + os.Getenv("DB_OPTIONS")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
