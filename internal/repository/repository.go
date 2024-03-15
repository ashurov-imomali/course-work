package repository

import "gorm.io/gorm"

type Repository struct {
	Db *gorm.DB
}

func GetRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}
