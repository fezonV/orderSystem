package domain

import "testing"

func TestNewOrderItem(t *testing.T) {
	product, _ := NewProduct(3, "Майка", "Белая майка из хлопка", 1000.0)
	orderItem, err := NewOrderItem(3, *product, 10)

	if err != nil {
		t.Fatalf("ожидали nil error, получили %v", err)
	}

	if orderItem.ID != 3 {
		t.Fatalf("Ожидали id = 3, получили %v", orderItem.ID)
	}

	if orderItem.Name != "Майка" {
		t.Fatalf("Ожидали name = Майка, получили %v", orderItem.Name)
	}

	if orderItem.Price != 1000.0 {
		t.Fatalf("Ожидали price = 1000, получили %v", orderItem.Price)
	}

	if orderItem.Quantity != 10 {
		t.Fatalf("Ожидали quantity = 10, получили %v", orderItem.Quantity)
	}
}

func TestNewOrderItemWithInvalidQuantity(t *testing.T) {
	product, _ := NewProduct(3, "Майка", "Белая майка из хлопка", 1000.0)
	_, err := NewOrderItem(3, *product, 0)

	if err != ErrInvalidQuantity {
		t.Fatalf("ожидали ErrInvalidQuantity, получили %v", err)
	}
}
