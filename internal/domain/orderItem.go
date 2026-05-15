package domain

import "fmt"

type OrderItem struct {
	ID          int64
	Name        string
	Price       float64
	Description string
}

func CreateOrderItem(id int64, name string, price float64, description string) (*OrderItem, error) {
	if id < 0 {
		return nil, fmt.Errorf("ID не может быть отрицательным")
	}
	if price < 0 {
		return nil, fmt.Errorf("Цена не может быть отрицательной")
	}
	return &OrderItem{
		ID:          id,
		Name:        name,
		Price:       price,
		Description: description,
	}, nil
}
