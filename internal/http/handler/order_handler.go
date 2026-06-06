package handler

import (
	"encoding/json"
	"net/http"
	"orderSystem/internal/usecase"
)

type OrderHandler struct {
	service usecase.OrderService
}

func NewOrderHandler(service usecase.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

// создать заказ
func (oh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order, err := oh.service.CreateOrder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := OrderResponse{
		OrderID:    order.ID,
		Status:     string(order.Status),
		OrderItems: []OrderItemResponse{},
		TotalSum:   order.TotalSum(),
	}
	json.NewEncoder(w).Encode(response)
}

// получить заказ

// добавить продукт в заказ

// оплатить заказ

// отменить заказ
