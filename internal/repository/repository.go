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

func (r *Repository) SelectServices(typeId, id int64, name string) ([]models.Service, error) {
	var srvs []models.Service
	query := r.Db.Where("type_id = ?", typeId)
	if id > 0 {
		query = query.Where("id = ?", id)
	}
	if name != "" {
		query = query.Where("name like ?", "%"+name+"%")
	}
	err := query.Find(&srvs).Error
	if err != nil {
		return nil, err
	}
	return srvs, err
}

func (r *Repository) ServiceById(id int64) (*models.ServiceResponse, error) {
	var srv models.ServiceResponse
	if err := r.Db.Select("s.id, s.name, s.description, s.image, t.name as type_name, r.commission, r.term, r.min, r.max").
		Table("services s").
		Joins("join rules r on r.id = s.rule_id").
		Joins("join types t on t.id = s.type_id").
		Where("s.id = ?", id).
		First(&srv).Error; err != nil {
		return nil, err
	}
	return &srv, nil
}

func (r *Repository) GetClientByLogin(login string) (*models.Client, error) {
	var client models.Client
	if err := r.Db.Where("login = ?", login).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *Repository) CreateClient(client *models.Client) error {
	return r.Db.Save(&client).Error
}
