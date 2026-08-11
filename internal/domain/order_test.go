package domain

import "testing"

func TestNewOrder(t *testing.T) {
	order, err := NewOrder(1)

	if err != nil {
		t.Fatalf("ожидали nil error, получили %v", err)
	}

	if order.ID() != 1 {
		t.Fatalf("ожидали id = 1, получили %v", order.ID())
	}

	if order.status != OrderStatusCreated {
		t.Fatalf("ожидали status = создан, получили %v", order.Status())
	}

	if len(order.Items()) != 0 {
		t.Fatalf("ожидали пустой список товаров, получили %v", len(order.Items()))
	}
}

func TestNewOrderWithNegativeID(t *testing.T) {
	_, err := NewOrder(-1)

	if err != ErrInvalidID {
		t.Fatalf("ожидали ErrInvalidID, получили %v", err)
	}
}

func TestAddProductToOrder(t *testing.T) {
	order, _ := NewOrder(1)
	product, _ := NewProduct(3, "Майка", "Белая майка из хлопка", 1000.0)

	err := order.AddProduct(*product, 2)

	if err != nil {
		t.Fatalf("ожидали nil error, получили %v", err)
	}

	if len(order.Items()) != 1 {
		t.Fatalf("ожидали 1 позицию, получили %v", len(order.Items()))
	}

	if order.Items()[0].ID() != 1 {
		t.Fatalf("ожидали item id = 1, получили %v", order.Items()[0].ID())
	}

	if order.Items()[0].Quantity() != 2 {
		t.Fatalf("ожидали quantity = 2, получили %v", order.Items()[0].Quantity())
	}
}

func TestTotalSum(t *testing.T) {
	order, _ := NewOrder(1)
	product1, _ := NewProduct(1, "Майка", "Белая майка", 1000.0)
	product2, _ := NewProduct(2, "Кроссовки", "Белые кроссовки", 7000.0)

	_ = order.AddProduct(*product1, 2)
	_ = order.AddProduct(*product2, 1)

	sum := order.TotalSum()

	if sum != 9000.0 {
		t.Fatalf("ожидали сумму 9000, получили %v", sum)
	}
}

func TestPayOrder(t *testing.T) {
	order, _ := NewOrder(1)
	product, _ := NewProduct(1, "Майка", "Белая майка", 1000.0)
	_ = order.AddProduct(*product, 2)

	err := order.Pay()

	if err != nil {
		t.Fatalf("ожидали nil error, получили %v", err)
	}

	if order.status != OrderStatusPaid {
		t.Fatalf("ожидали status = оплачен, получили %v", order.Status())
	}
}

func TestPayEmptyOrder(t *testing.T) {
	order, _ := NewOrder(1)

	err := order.Pay()

	if err != ErrOrderIsEmpty {
		t.Fatalf("ожидали ErrOrderIsEmpty, получили %v", err)
	}
}

func TestCancelOrder(t *testing.T) {
	order, _ := NewOrder(1)

	err := order.Cancel()

	if err != nil {
		t.Fatalf("ожидали nil error, получили %v", err)
	}

	if order.status != OrderStatusCanceled {
		t.Fatalf("ожидали status = отменен, получили %v", order.Status())
	}
}

func TestAddProductToPaidOrder(t *testing.T) {
	order, _ := NewOrder(1)
	product, _ := NewProduct(1, "Майка", "Белая майка", 1000.0)
	_ = order.AddProduct(*product, 2)
	_ = order.Pay()

	err := order.AddProduct(*product, 1)

	if err != ErrOrderAlreadyPaid {
		t.Fatalf("ожидали ErrOrderAlreadyPaid, получили %v", err)
	}
}

func TestAddProductToCanceledOrder(t *testing.T) {
	order, _ := NewOrder(1)
	product, _ := NewProduct(1, "Майка", "Белая майка", 1000.0)
	_ = order.Cancel()

	err := order.AddProduct(*product, 1)

	if err != ErrOrderAlreadyCanceled {
		t.Fatalf("ожидали ErrOrderAlreadyCanceled, получили %v", err)
	}
}
