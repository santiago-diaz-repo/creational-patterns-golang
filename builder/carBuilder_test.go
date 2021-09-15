package builder

import (
	"fmt"
	"testing"
)

func Test_Builder(t *testing.T) {
	car := carBuilder{}.Color("red").Brand("something").Build()
	fmt.Println(car)

	car2 := carBuilder{}.Color("red").Type("stuff").Build()
	fmt.Println(car2)
}
