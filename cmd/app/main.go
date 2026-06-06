package main

import (
	"fmt"
	"net/http"
	"orderSystem/internal/http/handler"
	"orderSystem/internal/storage/memory"
	"orderSystem/internal/usecase"
)

func main() {
	repo := memory.NewOrderRepository()
	service := usecase.NewOrderService(repo)
	orderHandler := handler.NewOrderHandler(*service)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /orders", orderHandler.CreateOrderHandler)
	mux.HandleFunc("GET /orders/{id}", orderHandler.GetOrderHandler)
	mux.HandleFunc("POST /orders/{id}/products", orderHandler.AddProductHandler)
	mux.HandleFunc("PATCH /orders/{id}/pay", orderHandler.OrderPayHandler)
	mux.HandleFunc("PATCH /orders/{id}/cancel", orderHandler.CancelOrderHandler)

	fmt.Println("сервер запущен на http://localhost:8080")

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}
