package database

import (
    "gorm.io/gorm"
)

type DatabaseHandle struct {
    Connection * gorm.DB 
}
