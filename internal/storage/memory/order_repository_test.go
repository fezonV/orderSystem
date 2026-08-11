package memory

import (
	"orderSystem/internal/domain"
	"testing"
)

func TestSaveOrder(t *testing.T) {
	or := NewOrderRepository()
	order, err := domain.NewOrder(1)

	if err != nil {
		t.Fatalf("не удалось создать заказ")
	}

	if order.ID() != 1 {
		t.Fatalf("ожидали id = 1, получили %v", order.ID())
	}

	err = or.Save(order)

	if err != nil {
		t.Fatalf("не удалось добавить заказ в репозиторий")
	}

	newOrder, err := or.GetByID(1)
	if err != nil {
		t.Fatalf("не удалось получить заказ из репозитория")
	}
	if newOrder.ID() != 1 {
		t.Fatalf("Не удалось получить заказ по id: ожидали id = 1, получили %v", newOrder.ID())
	}
}

func TestGetOrderByIDNotFound(t *testing.T) {
	or := NewOrderRepository()

	_, err := or.GetByID(1)

	if err == nil {
		t.Fatalf("получили заказ по неправильному id")
	}
}
