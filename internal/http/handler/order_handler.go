package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"orderSystem/internal/domain"
	"orderSystem/internal/usecase"
	"strconv"
)

type OrderHandler struct {
	service usecase.OrderService
}

func NewOrderHandler(service usecase.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func writeOrderError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	switch {
	case errors.Is(err, domain.ErrOrderNotFound):
		status = http.StatusNotFound
	case errors.Is(err, domain.ErrOrderAlreadyPaid),
		errors.Is(err, domain.ErrOrderAlreadyCanceled):
		status = http.StatusConflict
	case errors.Is(err, domain.ErrOrderIsEmpty),
		errors.Is(err, domain.ErrInvalidQuantity),
		errors.Is(err, domain.ErrInvalidPrice),
		errors.Is(err, domain.ErrInvalidID):
		status = http.StatusBadRequest
	}

	http.Error(w, err.Error(), status)
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
func (oh *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	order, err := oh.service.CreateOrder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := toOrderResponse(order)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// получить заказ
func (oh *OrderHandler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order, err := oh.service.GetOrder(id)
	if err != nil {
		writeOrderError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := toOrderResponse(order)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// добавить продукт в заказ
func (oh *OrderHandler) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var request AddProductToOrderRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := domain.NewProduct(request.ProductID, request.Name, request.Description, request.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order, err := oh.service.GetOrder(id)
	if err != nil {
		writeOrderError(w, err)
		return
	}
	err = oh.service.AddProductToOrder(order.ID, *product, int(request.Quantity))
	if err != nil {
		writeOrderError(w, err)
		return
	}
	updatedOrder, err := oh.service.GetOrder(order.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := toOrderResponse(updatedOrder)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// оплатить заказ
func (oh *OrderHandler) OrderPayHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := oh.service.GetOrder(id)
	if err != nil {
		writeOrderError(w, err)
		return
	}

	err = oh.service.PayOrder(order.ID)
	if err != nil {
		writeOrderError(w, err)
		return
	}

	updatedOrder, err := oh.service.GetOrder(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := toOrderResponse(updatedOrder)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// отменить заказ
func (oh *OrderHandler) CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = oh.service.CancelOrder(id)
	if err != nil {
		writeOrderError(w, err)
		return
	}
	updatedOrder, err := oh.service.GetOrder(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := toOrderResponse(updatedOrder)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
