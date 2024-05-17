package service

import "back-end/pkg/models"

type Service struct {
	Repo ReposMeths
}

func GetService(repo ReposMeths) SrvMeths {
	return &Service{Repo: repo}
}

func (s *Service) GetServices(typeId int64) ([]models.Service, error) {
	return s.Repo.GetService(typeId)
}
