package builder

type CarBuilder interface {
	Color(color string) CarBuilder
	Brand(brand string) CarBuilder
	Type(typeCa string) CarBuilder
	Price(price float64) CarBuilder
	Build() Car
}

type carBuilder struct {
	car
}

func (cb carBuilder) Build() Car {
	return car{
			color: cb.car.color,
			brand: cb.car.brand,
			typeC: cb.car.typeC,
			price: cb.car.price,
	}
}

func (cb carBuilder) Color(color string) CarBuilder {
	cb.car.color = color
	return cb
}

func (cb carBuilder) Brand(brand string) CarBuilder{
	cb.car.brand = brand
	return cb
}

func (cb carBuilder) Type(typeCa string) CarBuilder {
	cb.car.typeC = typeCa
	return cb
}

func (cb carBuilder) Price(price float64) CarBuilder {
	cb.car.price = price
	return cb
}