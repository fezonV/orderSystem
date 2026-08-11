package usecase

import (
	"orderSystem/internal/domain"
	"orderSystem/internal/storage/memory"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	os := memory.NewOrderRepository()
	service := NewOrderService(os)

	_, err := service.CreateOrder()
	if err != nil {
		t.Fatalf("ошибка при создании заказа: %v", err)
	}

	order, err := service.orderRepo.GetByID(1)

	if err != nil {
		t.Fatalf("ошибка при получении заказа: %v", err)
	}

	if order.ID() != 1 {
		t.Fatalf("ожидали id = 1, получили %v", order.ID())
	}
}

func TestAddProductToOrder(t *testing.T) {
	os := memory.NewOrderRepository()
	service := NewOrderService(os)
	_, err := service.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	pr, err := domain.NewProduct(123, "Чипсы", "чипсы картофельные со вкусом краба", 120.0)

	if err != nil {
		t.Fatalf("не удалось создать продукт для добавления в заказ: %v", err)
	}
	err = service.AddProductToOrder(1, *pr, 64)
	if err != nil {
		t.Fatalf("не удалось добавить продукт в заказ: %v", err)
	}

	order, err := service.orderRepo.GetByID(1)
	if err != nil {
		t.Fatalf("не удалось получить заказ: %v", err)
	}
	if len(order.Items()) != 1 {
		t.Fatalf("ожидали 1 позицию, получили %v", len(order.Items()))
	}
	if order.Items()[0].ProductID != 123 {
		t.Fatalf("ожидали product id = 123, получили %v", order.Items()[0].ProductID)
	}

	if order.Items()[0].ID != 1 {
		t.Fatalf("ожидали id = 1, получили %v", order.Items()[0].ID)
	}

	if order.Items()[0].Name != "Чипсы" {
		t.Fatalf("ожидали Name = Чипсы, получили %v", order.Items()[0].Name)
	}
	if order.Items()[0].Price != 120.0 {
		t.Fatalf("ожидали price = 120.0, получили %v", order.Items()[0].Price)
	}
}

func TestAddProductToOrderNotFound(t *testing.T) {
	repo := memory.NewOrderRepository()
	os := NewOrderService(repo)
	pr, err := domain.NewProduct(1, "кеды", "кеды жоские", 125.00)
	if err != nil {
		t.Fatalf("не удалось создать продукт: %v", err)
	}
	err = os.AddProductToOrder(1, *pr, 25)
	if err != domain.ErrOrderNotFound {
		t.Fatalf("ожидали ErrOrderNotFound, получили %v", err)
	}
}

func TestAddProductToOrderWithInvalidQuantity(t *testing.T) {
	repo := memory.NewOrderRepository()
	os := NewOrderService(repo)

	pr, err := domain.NewProduct(1, "рыба", "вяленая", 125.00)

	if err != nil {
		t.Fatalf("не удалось создать продукт: %v", err)
	}

	_, err = os.CreateOrder()

	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}
	err = os.AddProductToOrder(1, *pr, -2)

	if err != domain.ErrInvalidQuantity {
		t.Fatalf("ожидали ErrInvalidQuantity, получили %v", err)
	}
}

func TestGetOrder(t *testing.T) {
	repo := memory.NewOrderRepository()
	os := NewOrderService(repo)

	_, err := os.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ")
	}

	order, err := os.GetOrder(1)
	if err != nil {
		t.Fatalf("не удалось получить заказ")
	}

	if order.ID() != 1 {
		t.Fatalf("ожидали id = 1, получили %v", 1)
	}
}

func TestPayOrder(t *testing.T) {
	repo := memory.NewOrderRepository()
	os := NewOrderService(repo)

	_, err := os.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ")
	}
	pr, _ := domain.NewProduct(1, "рыба", "вяленая", 125.00)
	err = os.AddProductToOrder(1, *pr, 2)
	if err != nil {
		t.Fatalf("не удалось добавить продукт: %v", err)
	}

	err = os.PayOrder(1)
	if err != nil {
		t.Fatalf("не удалось оплатить заказ: %v", err)
	}
}

func TestCancelOrder(t *testing.T) {
	repo := memory.NewOrderRepository()
	os := NewOrderService(repo)

	_, err := os.CreateOrder()
	if err != nil {
		t.Fatalf("не удалось создать заказ")
	}
	os.CancelOrder(1)
	order, err := os.GetOrder(1)

	if order.Status() != domain.OrderStatusCanceled {
		t.Fatalf("ожидали отменен, получили %v", order.Status())
	}
}
