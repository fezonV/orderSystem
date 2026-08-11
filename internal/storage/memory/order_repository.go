package memory

import (
	"orderSystem/internal/domain"
	"orderSystem/internal/repository"
)

type OrderRepository struct {
	orders map[int64]*domain.Order
}

func NewOrderRepository() repository.OrderRepository {
	return &OrderRepository{
		orders: make(map[int64]*domain.Order),
	}
}
func (r *OrderRepository) Save(order *domain.Order) error {
	r.orders[order.ID()] = order
	return nil
}

func (r *OrderRepository) GetByID(id int64) (*domain.Order, error) {
	order, ok := r.orders[id]
	if !ok {
		return nil, domain.ErrOrderNotFound
	}

	return order, nil
}
