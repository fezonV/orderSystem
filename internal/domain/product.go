package domain

type Product struct {
	id          int64
	name        string
	description string
	price       Money
}

func (p Product) ID() int64 {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Description() string {
	return p.description
}

func (p Product) Price() Money {
	return p.price
}

func (p Product) validate() error {
	if p.id <= 0 {
		return ErrInvalidID
	}

	if p.price <= 0 {
		return ErrInvalidPrice
	}
	return nil
}

func NewProduct(id int64, name string, desc string, price Money) (*Product, error) {
	product := Product{
		id:          id,
		name:        name,
		description: desc,
		price:       price,
	}

	if err := product.validate(); err != nil {
		return nil, err
	}

	return &product, nil
}
