package factory

import "fmt"

type ConcreteProductTwo struct {
}

func NewConcreteProductTwo() Product {
	return ConcreteProductTwo{}
}

func (cpt ConcreteProductTwo) ExecuteStuffProduct() {
	fmt.Println("Invoking ConcreteProductTwo's executeStuffProduct method")
}
