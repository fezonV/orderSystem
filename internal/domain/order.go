package domain

import (
	"fmt"
)

type OrderStatus string

const (
	OrderStatusCreated  OrderStatus = "создан"
	OrderStatusPaid     OrderStatus = "оплачен"
	OrderStatusCanceled OrderStatus = "отменен"
)

type Order struct {
	ID         int64
	OrderItems map[int64]OrderItem
	Status     OrderStatus
	nextItemID int64
}

func NewOrder(id int64) (*Order, error) {
	if id < 0 {
		return nil, fmt.Errorf("Id не может быть отрицательным числом")
	}

	return &Order{
		ID:         id,
		OrderItems: make(map[int64]OrderItem),
		Status:     "создан",
		nextItemID: 0,
	}, nil
}

func (o *Order) AddOrderItem(oi OrderItem) {
	o.nextItemID += 1
	o.OrderItems[o.nextItemID] = oi
}
