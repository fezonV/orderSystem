package domain

import "testing"

func TestNewOrderItem(t *testing.T) {
	product, _ := NewProduct(3, "Майка", "Белая майка из хлопка", 1000.0)
	orderItem, err := NewOrderItem(3, *product, 10)

	if err != nil {
		t.Fatalf("ожидали nil error, получили %v", err)
	}

	if orderItem.ID != 3 {
		t.Fatalf("Ожидали id = 3, получили %v", product.ID)
	}

	if orderItem.Name != "Майка" {
		t.Fatalf("Ожидали name = Майка, получили %v", product.Name)
	}

	if orderItem.Price != 1000.0 {
		t.Fatalf("Ожидали description = Белая майка из хлопка, получили %v", product.Description)
	}

}
