package domain

import "fmt"

type OrderItem struct {
	Name        string
	Price       float64
	Description string
	Quantity    int
}

func NewOrderItem(name string, price float64, description string, quantity int) (*OrderItem, error) {
	if price < 0 {
		return nil, fmt.Errorf("цена не может быть отрицательной")
	}

	if quantity <= 0 {
		return nil, fmt.Errorf("количество должно быть положительным")
	}

	return &OrderItem{
		Name:        name,
		Price:       price,
		Description: description,
		Quantity:    quantity,
	}, nil
}
