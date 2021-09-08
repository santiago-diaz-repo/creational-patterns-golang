package factory

type ConcreteCreatorOne struct {
	Abstraction
}

func NewConcreteCreatorOne() Factory {
	return ConcreteCreatorOne{}
}

func (c ConcreteCreatorOne) CreateProduct() Product {
	p := NewConcreteProductOne()
	p.ExecuteStuffProduct()
	return p
}
