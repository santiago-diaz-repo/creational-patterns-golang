package builder

type Car interface {
	GetColor() string
	GetBrand() string
}

type car struct {
	price float64
	color string
	brand string
	typeC string
}

func (c car) GetColor() string {
	return c.color
}

func (c car) GetBrand() string {
	return c.brand
}