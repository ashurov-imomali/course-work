package service

import "back-end/pkg/models"

type ReposMeths interface {
	GetService(typeId int64) ([]models.Service, error)
}

type SrvMeths interface {
	GetServices(typeID int64) ([]models.Service, error)
}
