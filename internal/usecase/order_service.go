package usecase

import (
	"orderSystem/internal/domain"
	"orderSystem/internal/repository"
)

type OrderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(or *repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepo: *or,
	}
}

func (s *OrderService) CreateOrder(id int64) (*domain.Order, error) {
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
	quantity int) error {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return err
	}
	if err := order.AddProduct(p, quantity); err != nil {
		return err
	}
	return s.orderRepo.Save(order)

}
