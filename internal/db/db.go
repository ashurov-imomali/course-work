package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Database struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}

func GetConnection(dbStruct *Database) (*gorm.DB, error) {
	dbSettings := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbStruct.Host, dbStruct.Port, dbStruct.Username, dbStruct.Password, dbStruct.DatabaseName)
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             6 * time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(postgres.Open(dbSettings), &gorm.Config{Logger: dbLogger})
	log.New(os.Stdout, "2", log.LstdFlags)
	if err != nil {
		return nil, err
	}
	return db, nil
}
