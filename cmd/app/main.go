package main

import (
	"fmt"
	"orderSystem/internal/domain"
)

func main() {

	Orders, _ := domain.NewOrder(123)

	OrderItem1, err := domain.NewOrderItem("Кроссы", 7000, "NIKE")
	if err != nil {
		panic(err)
	}
	OrderItem2, err := domain.NewOrderItem("Мяч", 3000, "хз")
	if err != nil {
		panic(err)
	}

	OrderItem3, err := domain.NewOrderItem("Сиги", 200, "мальборо красный в моем кармане")
	if err != nil {
		panic(err)
	}
	Orders.AddOrderItem(*OrderItem1)
	Orders.AddOrderItem(*OrderItem2)
	Orders.AddOrderItem(*OrderItem3)

	fmt.Println(Orders)

}
