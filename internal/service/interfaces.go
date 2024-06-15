package service

import "back-end/pkg/models"

type ReposMeths interface {
	SelectServices(typeId, id int64, name string) ([]models.Service, error)
	ServiceById(id int64) (*models.ServiceResponse, error)
	GetClientByLogin(login string) (*models.Client, error)
	CreateClient(client *models.Client) error
	CreateCreditSrv(req *models.SrvReq, typ int64, s string) error
	GetServiceById(id int64) (*models.Service, error)
}

type SrvMeths interface {
	GetServices(typeId, id int64, name string) ([]models.Service, error)
	GetService(id int64) (*models.ServiceResponse, error)
	Registration(newClient *models.Client) *Error
	Login(login *models.Login) (string, *Error)
	VerifyLogin(otp *models.VerifyLogin) *Error
	CreateSrvForClient(srv *models.SrvReq) *Error
}
