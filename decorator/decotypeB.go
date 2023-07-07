package decorator

import (
	"fmt"
)

type DecoratorTypeB struct {
	Product ProductA
	somethingB string
}

func (dtb *DecoratorTypeB)addSomething()  {
	dtb.somethingB = "Add something b"
}


func (dtb *DecoratorTypeB)GetProduct() {
	dtb.Product.PrintAnything()
	dtb.addSomething()
	fmt.Println()
	fmt.Printf("Action log: %s",dtb.somethingB)
}