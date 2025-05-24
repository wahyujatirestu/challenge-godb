package service

import (
	"challenge-godb/entity"
	"challenge-godb/repository"
	"errors"
)

type CustomerService interface {
	CreateCustomer(c entity.Customer) error
	GetAllCustomers()([]entity.Customer, error)
	GetCustomerById(id int)(entity.Customer, error)
	UpdateCustomer(c entity.Customer) error
	DeleteCustomer(id int) error
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(r repository.CustomerRepository) CustomerService {
	return &customerService{r}
}

func (s *customerService) CreateCustomer(c entity.Customer) error {
	if s.repo.IsIdExist(c.CustomerId) {
		return errors.New("customer ID already exist")
	}
	return s.repo.Create(c)
}

func (s *customerService) GetAllCustomers() ([]entity.Customer, error) {
	return s.repo.FindAll()
}

func (s *customerService) GetCustomerById(id int) (entity.Customer, error) {
	return s.repo.FindById(id)
}

func (s *customerService) UpdateCustomer(c entity.Customer) error {
	return s.repo.Update(c)
}

func (s *customerService) DeleteCustomer(id int) error {
	return s.repo.Delete(id)
}