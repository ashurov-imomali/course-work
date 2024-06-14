package service

import "back-end/pkg/models"

type ReposMeths interface {
	SelectServices(typeId, id int64, name string) ([]models.Service, error)
	ServiceById(id int64) (*models.ServiceResponse, error)
	GetClientByLogin(login string) (*models.Client, error)
	CreateClient(client *models.Client) error
}

type SrvMeths interface {
	GetServices(typeId, id int64, name string) ([]models.Service, error)
	GetService(id int64) (*models.ServiceResponse, error)
	Registration(newClient *models.Client) *Error
	Login(login *models.Login) (string, *Error)
	VerifyLogin(otp *models.VerifyLogin) *Error
}
