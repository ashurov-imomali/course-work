package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}

func GetConnection(dbStruct *Database) (*gorm.DB, error) {
	dbSettings := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s")
	db, err := gorm.Open(postgres.Open(dbSettings), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
