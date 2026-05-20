package main

import (
	"fmt"
	"orderSystem/internal/domain"
)

func main() {
	Sneakers, err := domain.NewProduct(1, "ботинки", "спортивные ботинки", 5000)
	if err != nil {
		panic(err)
	}
	Cigs, err := domain.NewProduct(1, "сиги", "мальборо", 250)

	SneakersOrder, err := domain.NewOrderItem(*Sneakers, 2)
	if err != nil {
		panic(err)
	}

	CigsOrder, err := domain.NewOrderItem(*Cigs, 10)
	if err != nil {
		panic(err)
	}

	Order1, err := domain.NewOrder(1)
	if err != nil {
		panic(err)
	}

	Order1.AddOrderItem(*CigsOrder)
	Order1.AddOrderItem(*SneakersOrder)

	for _, v := range Order1.OrderItems {
		fmt.Println("Позиция: ", v.Name, "Количество: ", v.Quantity)
	}
}
