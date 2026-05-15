package domain

import "fmt"

type OrderStatus string

const (
	OrderStatusCreated  OrderStatus = "создан"
	OrderStatusPaid     OrderStatus = "оплачен"
	OrderStatusCanceled OrderStatus = "отменен"
)

type Order struct {
	ID         int64
	OrderItems []OrderItem
	Status     OrderStatus
}

func CreateOrder(id int64, status OrderStatus) (*Order, error) {
	if id < 0 {
		return nil, fmt.Errorf("Id не может быть отрицательным числом")
	}

	return &Order{
		ID:         id,
		OrderItems: make([]OrderItem, 0),
		Status:     "создан",
	}, nil
}
