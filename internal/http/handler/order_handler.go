package handler

import (
	"encoding/json"
	"net/http"
	"orderSystem/internal/domain"
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
func toOrderResponse(order *domain.Order) OrderResponse {
	items := make([]OrderItemResponse, 0, len(order.OrderItems))

	for _, item := range order.OrderItems {
		items = append(items, OrderItemResponse{
			ProductID: item.ProductID,
			Name:      item.Name,
			Price:     item.Price,
			Quantity:  int64(item.Quantity),
		})
	}

	return OrderResponse{
		OrderID:    order.ID,
		Status:     string(order.Status),
		OrderItems: items,
		TotalSum:   order.TotalSum(),
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
	response := toOrderResponse(order)
	json.NewEncoder(w).Encode(response)
}

// получить заказ

// добавить продукт в заказ

// оплатить заказ

// отменить заказ
