package factory

import "fmt"

type ConcreteProductOne struct {
}

func NewConcreteProductOne() Product {
	return ConcreteProductOne{}
}

func (cpo ConcreteProductOne) ExecuteStuffProduct() {
	fmt.Println("Invoking ConcreteProductOne's executeStuffProduct method")
}
