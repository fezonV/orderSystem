package handler

type OrderItemResponse struct {
	ProductID    int64  `json:"product_id"`
	Name         string `json:"name"`
	PriceKopecks int64  `json:"price_kopecks"`
	Quantity     int64  `json:"quantity"`
}

type OrderResponse struct {
	OrderID         int64               `json:"order_id"`
	Status          string              `json:"status"`
	OrderItems      []OrderItemResponse `json:"orderItems"`
	TotalSumKopecks int64               `json:"total_sum_kopecks"`
}

type AddProductToOrderRequest struct {
	ProductID    int64  `json:"product_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	PriceKopecks int64  `json:"price_kopecks"`
	Quantity     int64  `json:"quantity"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
