package domain

type OrderStatus string

const (
	OrderStatusCreated  OrderStatus = "создан"
	OrderStatusPaid     OrderStatus = "оплачен"
	OrderStatusCanceled OrderStatus = "отменен"
)

type Order struct {
	id         int64
	items      []OrderItem
	status     OrderStatus
	nextItemID int64
}

// копирование заказа без ссылок
func (o Order) Clone() Order {
	result := o

	result.items = make([]OrderItem, len(o.items))
	copy(result.items, o.items)

	return result
}

func (o *Order) ID() int64 {
	return o.id
}

func (o *Order) Status() OrderStatus {
	return o.status
}

// Возвращает копию списка элементов заказа
func (o *Order) Items() []OrderItem {
	items := make([]OrderItem, len(o.items))

	copy(items, o.items)

	return items
}

func NewOrder(id int64) (*Order, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}

	return &Order{
		id:     id,
		items:  make([]OrderItem, 0),
		status: OrderStatusCreated,
	}, nil
}

func (o *Order) AddProduct(product Product, quantity int) error {
	if o.status == OrderStatusCanceled {
		return ErrOrderAlreadyCanceled
	}
	if o.status == OrderStatusPaid {
		return ErrOrderAlreadyPaid
	}
	itemID := o.nextItemID + 1
	oi, err := NewOrderItem(itemID, product, quantity)
	if err != nil {
		return err
	}

	o.nextItemID = itemID
	o.items = append(o.items, *oi)
	return nil
}

func (o Order) TotalSum() Money {
	var sum Money
	for _, v := range o.items {
		sum += v.Price() * Money(v.Quantity())
	}
	return sum
}

func (o *Order) Pay() error {
	if o.status == OrderStatusPaid {
		return ErrOrderAlreadyPaid
	}
	if o.status == OrderStatusCanceled {
		return ErrOrderAlreadyCanceled
	}
	if len(o.items) == 0 {
		return ErrOrderIsEmpty
	}
	o.status = OrderStatusPaid
	return nil
}

// TODO
// Изменить статус на отмемнен
func (o *Order) Cancel() error {
	if o.Status() == OrderStatusPaid {
		return ErrOrderAlreadyPaid
	}
	if o.Status() == OrderStatusCanceled {
		return ErrOrderAlreadyCanceled
	}
	o.status = OrderStatusCanceled
	return nil
}
