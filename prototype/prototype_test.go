package prototype

import (
	"fmt"
	"testing"
)

func Test_Prototype(t *testing.T) {
	first := concretePrototype{
		attr1: "A",
		attr2: "B",
		attr3: "C",
	}

	first.print("")

	var second *concretePrototype = first.clone().(*concretePrototype)

	third := first.clone().(*concretePrototype)

	second.print("")
	third.print("")

	first.setAttr1("NEW")
	fmt.Println("after")
	first.print("")

	second.print("")
	third.print("")

}
