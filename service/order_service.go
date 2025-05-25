package service

import (
	"challenge-godb/entity"
	"challenge-godb/repository"
	"errors"
	"time"
)

type OrderService interface {
	CreateOrder(order entity.Order, details []entity.OrderDetail) error
	GetAllOrder() ([]entity.Order, error)
	GetOrderById(orderId int) (*entity.Order, []entity.OrderDetail, error)
	CompleteOrder(orderId int, completionDate string) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(r repository.OrderRepository) OrderService {
	return &orderService{r}
}

func (s *orderService) CreateOrder(order entity.Order, details []entity.OrderDetail) error {
	if s.repo.IsOrderIdExist(order.Order_Id) {
		return errors.New("Order ID already exists. Please enter a different ID.")
	}
	if !s.repo.IsCustomerIdExist(order.Customer_Id) {
		return errors.New("Customer not found.")
	}
	order.Order_Date = time.Now()
	return s.repo.Create(order, details)
}

func (s *orderService) GetAllOrder() ([]entity.Order, error) {
	return s.repo.GetAll()
}

func (s *orderService) GetOrderById(orderId int) (*entity.Order, []entity.OrderDetail, error) {
	if !s.repo.IsOrderIdExist(orderId) {
		return nil, nil, errors.New("Order not found.")
	}
	return s.repo.FindById(orderId)
}

func (s *orderService) CompleteOrder(orderId int, completionDate string) error {
	if !s.repo.IsOrderIdExist(orderId) {
		return errors.New("Order not found.")
	}
	return s.repo.CompleteOrder(orderId, completionDate)
}
