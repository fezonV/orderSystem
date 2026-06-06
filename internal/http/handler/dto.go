package handler

type OrderItemResponse struct {
	OrderID   int64   `json:"order_id"`
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int64   `json:"quantity"`
}

type OrderResponse struct {
	OrderID    int64               `json:"order_id"`
	Status     string              `json:"status"`
	OrderItems []OrderItemResponse `json:"orderItems"`
	TotalSum   float64             `json:"total_sum"`
}

type AddProductToOrderRequest struct {
	ProductID   int64   `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
