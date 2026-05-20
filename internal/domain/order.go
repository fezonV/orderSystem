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
	OrderItems []OrderItem
	Status     OrderStatus
	nextItemID int64
}

func NewOrder(id int64) (*Order, error) {
	if id < 0 {
		return nil, fmt.Errorf("Id не может быть отрицательным числом")
	}

	return &Order{
		ID:         id,
		OrderItems: make([]OrderItem, 0),
		Status:     "создан",
	}, nil
}

func (o *Order) AddProduct(product Product, quantity int) error {
	oi, err := NewOrderItem(product, quantity)
	if err != nil {
		return err
	}
	o.nextItemID += 1
	o.OrderItems = append(o.OrderItems, *oi)
	return nil
}

// TODO
// сделать возможность хранить позиции заказа

//TODO
//Добавлять позицию

//TODO
// Считать итоговую сумму заказа

//TODO
// Изменить статус на оплачен

// TODO
// Изменить статус на отмемнен
