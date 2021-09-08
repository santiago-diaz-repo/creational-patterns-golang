package factory

import "testing"

func TestAbstraction_DoStuff(t *testing.T) {
	subject := NewConcreteCreatorOne()
	subject.DoStuff()
	subject.CreateProduct()
}
