package repository

import "gorm.io/gorm"

type Repository struct {
	Db *gorm.DB
}

func GetRepository() *Repository {
	return &Repository{}
}
