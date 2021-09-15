package prototype

import "fmt"

type concretePrototype struct {
	attr1 string
	attr2 string
	attr3 string
}

func (p *concretePrototype) print(string2 string) {
	fmt.Printf("%s , %s, %s \n", p.attr1, p.attr2, p.attr3)
}

func (p *concretePrototype) setAttr1(attr1 string) {
	p.attr1 = attr1
}

func (p *concretePrototype) clone() Prototyper {
	return &concretePrototype{
		attr1: p.attr1,
		attr2: p.attr2,
		attr3: p.attr3,
	}
}
