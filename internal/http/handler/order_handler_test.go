package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"orderSystem/internal/domain"
	"orderSystem/internal/storage/memory"
	"orderSystem/internal/usecase"
	"strings"
	"testing"
)

func TestCreateOrderHandler(t *testing.T) {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := NewOrderHandler(*service)

	request := httptest.NewRequest(http.MethodPost, "/orders", nil)
	responseRecorder := httptest.NewRecorder()

	orderHandler.CreateOrderHandler(responseRecorder, request)

	if responseRecorder.Code != http.StatusCreated {
		t.Fatalf("ожидали status = 201, получили %v", responseRecorder.Code)
	}

	var response OrderResponse
	err := json.NewDecoder(responseRecorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("не удалось прочитать ответ: %v", err)
	}

	if response.OrderID != 1 {
		t.Fatalf("ожидали order id = 1, получили %v", response.OrderID)
	}

	if response.Status != string(domain.OrderStatusCreated) {
		t.Fatalf("ожидали status = создан, получили %v", response.Status)
	}

	if len(response.OrderItems) != 0 {
		t.Fatalf("ожидали пустой список товаров, получили %v", len(response.OrderItems))
	}
}

func TestGetOrderHandler(t *testing.T) {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := NewOrderHandler(*service)

	_, err := service.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orders/{id}", orderHandler.GetOrderHandler)

	request := httptest.NewRequest(http.MethodGet, "/orders/1", nil)
	responseRecorder := httptest.NewRecorder()

	mux.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("ожидали status = 200, получили %v", responseRecorder.Code)
	}

	var response OrderResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("не удалось прочитать ответ: %v", err)
	}

	if response.OrderID != 1 {
		t.Fatalf("ожидали order id = 1, получили %v", response.OrderID)
	}
}

func TestGetOrderHandlerNotFound(t *testing.T) {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := NewOrderHandler(*service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orders/{id}", orderHandler.GetOrderHandler)

	request := httptest.NewRequest(http.MethodGet, "/orders/1", nil)
	responseRecorder := httptest.NewRecorder()

	mux.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusNotFound {
		t.Fatalf("ожидали status = 404, получили %v", responseRecorder.Code)
	}
}

func TestGetOrderHandlerWithInvalidID(t *testing.T) {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := NewOrderHandler(*service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orders/{id}", orderHandler.GetOrderHandler)

	request := httptest.NewRequest(http.MethodGet, "/orders/text", nil)
	responseRecorder := httptest.NewRecorder()

	mux.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("ожидали status = 400, получили %v", responseRecorder.Code)
	}
}

func TestAddProductHandler(t *testing.T) {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := NewOrderHandler(*service)

	_, err := service.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	body := strings.NewReader(`{
		"product_id": 3,
		"name": "Майка",
		"description": "Белая майка",
		"price_kopecks": 100000,
		"quantity": 2
	}`)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders/{id}/products", orderHandler.AddProductHandler)

	request := httptest.NewRequest(http.MethodPost, "/orders/1/products", body)
	responseRecorder := httptest.NewRecorder()

	mux.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("ожидали status = 200, получили %v", responseRecorder.Code)
	}

	var response OrderResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("не удалось прочитать ответ: %v", err)
	}

	if len(response.OrderItems) != 1 {
		t.Fatalf("ожидали 1 позицию, получили %v", len(response.OrderItems))
	}

	if response.OrderItems[0].ProductID != 3 {
		t.Fatalf("ожидали product id = 3, получили %v", response.OrderItems[0].ProductID)
	}

	if response.OrderItems[0].Quantity != 2 {
		t.Fatalf("ожидали quantity = 2, получили %v", response.OrderItems[0].Quantity)
	}

	if response.TotalSumKopecks != 200_000 {
		t.Fatalf("ожидали total sum = 200000 копеек, получили %v", response.TotalSumKopecks)
	}
}

func TestAddProductHandlerWithInvalidQuantity(t *testing.T) {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := NewOrderHandler(*service)

	_, err := service.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	body := strings.NewReader(`{
		"product_id": 3,
		"name": "Майка",
		"description": "Белая майка",
		"price_kopecks": 100000,
		"quantity": 0
	}`)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders/{id}/products", orderHandler.AddProductHandler)

	request := httptest.NewRequest(http.MethodPost, "/orders/1/products", body)
	responseRecorder := httptest.NewRecorder()

	mux.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("ожидали status = 400, получили %v", responseRecorder.Code)
	}
}

func TestOrderPayHandler(t *testing.T) {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := NewOrderHandler(*service)

	_, err := service.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	product, err := domain.NewProduct(3, "Майка", "Белая майка", 100_000)
	if err != nil {
		t.Fatalf("не удалось создать товар: %v", err)
	}

	_, err = service.AddProductToOrder(1, *product, 2)
	if err != nil {
		t.Fatalf("не удалось добавить товар в заказ: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /orders/{id}/pay", orderHandler.OrderPayHandler)

	request := httptest.NewRequest(http.MethodPatch, "/orders/1/pay", nil)
	responseRecorder := httptest.NewRecorder()

	mux.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("ожидали status = 200, получили %v", responseRecorder.Code)
	}

	var response OrderResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("не удалось прочитать ответ: %v", err)
	}

	if response.Status != string(domain.OrderStatusPaid) {
		t.Fatalf("ожидали status = оплачен, получили %v", response.Status)
	}
}

func TestOrderPayHandlerWithEmptyOrder(t *testing.T) {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := NewOrderHandler(*service)

	_, err := service.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /orders/{id}/pay", orderHandler.OrderPayHandler)

	request := httptest.NewRequest(http.MethodPatch, "/orders/1/pay", nil)
	responseRecorder := httptest.NewRecorder()

	mux.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("ожидали status = 400, получили %v", responseRecorder.Code)
	}
}

func TestCancelOrderHandler(t *testing.T) {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := NewOrderHandler(*service)

	_, err := service.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /orders/{id}/cancel", orderHandler.CancelOrderHandler)

	request := httptest.NewRequest(http.MethodPatch, "/orders/1/cancel", nil)
	responseRecorder := httptest.NewRecorder()

	mux.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("ожидали status = 200, получили %v", responseRecorder.Code)
	}

	var response OrderResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("не удалось прочитать ответ: %v", err)
	}

	if response.Status != string(domain.OrderStatusCanceled) {
		t.Fatalf("ожидали status = отменен, получили %v", response.Status)
	}
}

func TestCancelPaidOrderHandler(t *testing.T) {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := NewOrderHandler(*service)

	_, err := service.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	product, err := domain.NewProduct(3, "Майка", "Белая майка", 100_000)
	if err != nil {
		t.Fatalf("не удалось создать товар: %v", err)
	}

	_, err = service.AddProductToOrder(1, *product, 2)
	if err != nil {
		t.Fatalf("не удалось добавить товар в заказ: %v", err)
	}

	err = service.PayOrder(1)
	if err != nil {
		t.Fatalf("не удалось оплатить заказ: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /orders/{id}/cancel", orderHandler.CancelOrderHandler)

	request := httptest.NewRequest(http.MethodPatch, "/orders/1/cancel", nil)
	responseRecorder := httptest.NewRecorder()

	mux.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusConflict {
		t.Fatalf("ожидали status = 409, получили %v", responseRecorder.Code)
	}
}
