package main

import (
	"orderSystem/internal/domain"
	"orderSystem/internal/storage/memory"
	"orderSystem/internal/usecase"
)

func main() {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)

	product, err := domain.NewProduct(1, "Кроссовки", "Кроссовки белые", 7000.0)
	if err != nil {
		panic(err)
	}
	order, err := service.CreateOrder()
	err = service.AddProductToOrder(order.ID, *product, 2)
}
