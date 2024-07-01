package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(dburl string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dburl), &gorm.Config{})
	if err != nil {
		panic("Database: failed making connection using sqlite to database " + dburl)
	}
	return db
}

func GetConnection(dburl string) DatabaseHandle {
	connectionDB := Connect(dburl)
    return  DatabaseHandle{Connection: connectionDB}
}
