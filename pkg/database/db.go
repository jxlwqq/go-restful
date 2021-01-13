package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func New(dsn string, cfg *gorm.Config) (*DB, error) {
	 db, err := gorm.Open(mysql.Open(dsn), cfg)
	 return &DB{db}, err
}