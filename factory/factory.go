package factory

import "fmt"

type Factory interface {
	DoStuff()
	CreateProduct() Product
}

type Abstraction struct {
}

func (a Abstraction) DoStuff() {
	fmt.Println("doStuff")
}
