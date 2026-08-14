package repository

import "orderSystem/internal/domain"

type OrderRepository interface {
	Save(order *domain.Order) error
	GetByID(id int64) (domain.Order, error)
}
