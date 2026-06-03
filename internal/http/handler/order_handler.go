package handler

import "orderSystem/internal/usecase"

type OrderHandler struct {
	service usecase.OrderService
}

func NewOrderHandler(service usecase.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

// создать заказ

// получить заказ

// добавить продукт в заказ

// оплатить заказ

// отменить заказ
