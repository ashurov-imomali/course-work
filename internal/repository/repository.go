package repository

import (
	"back-end/internal/service"
	"back-end/pkg/models"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func GetRepository(db *gorm.DB) service.ReposMeths {
	return &Repository{db}
}

func (r *Repository) GetService(typeId int64) ([]models.Service, error) {
	var srvs []models.Service
	err := r.Db.Where("type_id = ?", typeId).Find(&srvs).Error
	if err != nil {
		return nil, err
	}
	return srvs, err
}
