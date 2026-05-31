package domain

import "testing"

func TestNewProduct(t *testing.T) {
	product, err := NewProduct(3, "Майка", "Белая майка из хлопка", 1000.0)

	if err != nil {
		t.Fatalf("ожидали nil error, получили %v", err)
	}

	if product.ID != 3 {
		t.Fatalf("Ожидали id = 3, получили %v", product.ID)
	}

	if product.Name != "Майка" {
		t.Fatalf("Ожидали name = Майка, получили %v", product.Name)
	}

	if product.Description != "Белая майка из хлопка" {
		t.Fatalf("Ожидали description = Белая майка из хлопка, получили %v", product.Description)
	}

	if product.Price != 1000.0 {
		t.Fatalf("Ожидали price = 1000, получили %v", product.Price)
	}
}

func TestNewProductWithNegativeID(t *testing.T) {
	_, err := NewProduct(-1, "Майка", "Белая майка из хлопка", 1000.0)

	if err != ErrInvalidID {
		t.Fatalf("ожидали ErrInvalidID, получили %v", err)
	}
}

func TestNewProductWithNegativePrice(t *testing.T) {
	_, err := NewProduct(3, "Майка", "Белая майка из хлопка", -1000.0)

	if err != ErrInvalidPrice {
		t.Fatalf("ожидали ErrInvalidPrice, получили %v", err)
	}
}
