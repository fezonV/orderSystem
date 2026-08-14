package usecase

import (
	"orderSystem/internal/domain"
	"orderSystem/internal/repository"
	"sync/atomic"
)

type OrderService struct {
	orderRepo   repository.OrderRepository
	nextOrderID int64
}

func NewOrderService(or repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepo: or,
	}
}

func (s *OrderService) CreateOrder() (*domain.Order, error) {
	id := atomic.AddInt64(&s.nextOrderID, 1)
	order, err := domain.NewOrder(id)
	if err != nil {
		return nil, err
	}
	if err := s.orderRepo.Save(order); err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) AddProductToOrder(
	orderID int64,
	p domain.Product,
	quantity int) (domain.Order, error) {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return domain.Order{}, err
	}
	if err := order.AddProduct(p, quantity); err != nil {
		return domain.Order{}, err
	}
	return order, s.orderRepo.Save(&order)

}

func (s *OrderService) GetOrder(id int64) (domain.Order, error) {
	return s.orderRepo.GetByID(id)
}

func (s *OrderService) PayOrder(id int64) error {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return err
	}
	if err := order.Pay(); err != nil {
		return err
	}
	return s.orderRepo.Save(&order)
}

func (s *OrderService) CancelOrder(id int64) error {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return err
	}
	if err := order.Cancel(); err != nil {
		return err
	}
	return s.orderRepo.Save(&order)
}
