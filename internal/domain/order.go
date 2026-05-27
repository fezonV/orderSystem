package domain

type OrderStatus string

const (
	OrderStatusCreated  OrderStatus = "создан"
	OrderStatusPaid     OrderStatus = "оплачен"
	OrderStatusCanceled OrderStatus = "отменен"
)

type Order struct {
	ID         int64
	OrderItems []OrderItem
	Status     OrderStatus
	nextItemID int64
}

func NewOrder(id int64) (*Order, error) {
	if id < 0 {
		return nil, ErrInvalidID
	}

	return &Order{
		ID:         id,
		OrderItems: make([]OrderItem, 0),
		Status:     OrderStatusCreated,
	}, nil
}

func (o *Order) AddProduct(product Product, quantity int) error {
	if o.Status == OrderStatusCanceled {
		return ErrOrderAlreadyCanceled
	}
	if o.Status == OrderStatusPaid {
		return ErrOrderAlreadyPaid
	}
	itemID := o.nextItemID + 1
	oi, err := NewOrderItem(itemID, product, quantity)
	if err != nil {
		return err
	}

	o.nextItemID = itemID
	o.OrderItems = append(o.OrderItems, *oi)
	return nil
}

// TODO
// Считать итоговую сумму заказа
func (o Order) TotalSum() float64 {
	sum := 0.0
	for _, v := range o.OrderItems {
		sum += v.Price * float64(v.Quantity)
	}
	return sum
}

// TODO
// Изменить статус на оплачен
func (o *Order) ChangeStatusPaid() error {
	if o.Status == OrderStatusPaid {
		return ErrOrderAlreadyPaid
	}
	if o.Status == OrderStatusCanceled {
		return ErrOrderAlreadyCanceled
	}
	if len(o.OrderItems) == 0 {
		return ErrOrderIsEmpty
	}
	o.Status = OrderStatusPaid
	return nil
}

// TODO
// Изменить статус на отмемнен
func (o *Order) ChangeStatusCanceled() error {
	if o.Status == OrderStatusPaid {
		return ErrOrderAlreadyPaid
	}
	if o.Status == OrderStatusCanceled {
		return ErrOrderAlreadyCanceled
	}
	o.Status = OrderStatusCanceled
	return nil
}
