package service

import (
	"challenge-godb/entity"
	"challenge-godb/repository"
	"errors"
)

type ServiceService interface {
	CreateService(s entity.Service) error
	GetAllService() ([]entity.Service, error)
	GetServiceById(id int) (entity.Service, error)
	UpdateService(s entity.Service) error
	DeleteService(id int) error
}

type serviceService struct {
	repo repository.ServiceRepository
}

func NewServiceService(r repository.ServiceRepository) ServiceService {
	return &serviceService{r}
}

func (s *serviceService) CreateService(service entity.Service) error {
	if s.repo.IsIdExist(service.Service_Id){
		return errors.New("Service ID already exist")
	}

	return s.repo.Create(service)
}

func (s *serviceService) GetAllService() ([]entity.Service, error) {
	return s.repo.GetAll()
}

func (s *serviceService) GetServiceById(id int) (entity.Service, error) {
	return s.repo.GetById(id)
}

func (s *serviceService) UpdateService(service entity.Service) error {
	return s.repo.Update(service)
}

func (s *serviceService) DeleteService(id int) error {
	return s.repo.Delete(id)
}