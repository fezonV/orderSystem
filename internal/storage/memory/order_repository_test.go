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

func TestSaveStoresCopy(t *testing.T) {
	or := NewOrderRepository()

	order, err := domain.NewOrder(1)
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	if err := or.Save(order); err != nil {
		t.Fatalf("не удалось сохранить заказ: %v", err)
	}

	// Изменение исходного заказа после Save не должно менять сохраненный заказ.
	if err := order.Cancel(); err != nil {
		t.Fatalf("не удалось отменить заказ: %v", err)
	}

	savedOrder, err := or.GetByID(1)
	if err != nil {
		t.Fatalf("не удалось получить заказ: %v", err)
	}

	if savedOrder.Status() != domain.OrderStatusCreated {
		t.Fatalf(
			"ожидали статус %v, получили %v",
			domain.OrderStatusCreated,
			savedOrder.Status(),
		)
	}
}

func TestGetByIDReturnsCopy(t *testing.T) {
	or := NewOrderRepository()

	order, err := domain.NewOrder(1)
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	if err := or.Save(order); err != nil {
		t.Fatalf("не удалось сохранить заказ: %v", err)
	}

	// Изменение результата GetByID без Save не должно менять сохраненный заказ.
	receivedOrder, err := or.GetByID(1)
	if err != nil {
		t.Fatalf("не удалось получить заказ: %v", err)
	}

	if err := receivedOrder.Cancel(); err != nil {
		t.Fatalf("не удалось отменить заказ: %v", err)
	}

	savedOrder, err := or.GetByID(1)
	if err != nil {
		t.Fatalf("не удалось повторно получить заказ: %v", err)
	}

	if savedOrder.Status() != domain.OrderStatusCreated {
		t.Fatalf(
			"ожидали статус %v, получили %v",
			domain.OrderStatusCreated,
			savedOrder.Status(),
		)
	}
}
