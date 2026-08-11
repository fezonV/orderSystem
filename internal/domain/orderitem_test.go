package domain

import "testing"

func TestNewOrderItem(t *testing.T) {
	product, _ := NewProduct(3, "Майка", "Белая майка из хлопка", 100_000)
	orderItem, err := NewOrderItem(3, *product, 10)

	if err != nil {
		t.Fatalf("ожидали nil error, получили %v", err)
	}

	if orderItem.ID() != 3 {
		t.Fatalf("Ожидали id = 3, получили %v", orderItem.ID())
	}

	if orderItem.Name() != "Майка" {
		t.Fatalf("Ожидали name = Майка, получили %v", orderItem.Name())
	}

	if orderItem.Price() != 100_000 {
		t.Fatalf("Ожидали price = 100000 копеек, получили %v", orderItem.Price())
	}

	if orderItem.Quantity() != 10 {
		t.Fatalf("Ожидали quantity = 10, получили %v", orderItem.Quantity())
	}
}

func TestNewOrderItemWithInvalidQuantity(t *testing.T) {
	product, _ := NewProduct(3, "Майка", "Белая майка из хлопка", 100_000)
	_, err := NewOrderItem(3, *product, 0)

	if err != ErrInvalidQuantity {
		t.Fatalf("ожидали ErrInvalidQuantity, получили %v", err)
	}
}

func TestNewOrderItemWithInvalidID(t *testing.T) {
	product, _ := NewProduct(1, "Майка", "Белая майка", 100_000)

	_, err := NewOrderItem(0, *product, 1)

	if err != ErrInvalidID {
		t.Fatalf("ожидали ErrInvalidID, получили %v", err)
	}
}

func TestNewOrderItemWithEmptyProduct(t *testing.T) {
	var product Product

	_, err := NewOrderItem(1, product, 1)

	if err != ErrInvalidID {
		t.Fatalf("ожидали ErrInvalidID для пустого Product, получили %v", err)
	}
}
