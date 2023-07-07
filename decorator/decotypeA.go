package decorator

import (
	"fmt"
)

type DecoratorTypeA struct {
	Product ProductA
	somethingA string
}

func (dta *DecoratorTypeA)addSomething()  {
	dta.somethingA = "Add something A"
}


func (dta *DecoratorTypeA)GetProduct() {
	dta.Product.PrintAnything()
	dta.addSomething()
	fmt.Println()
	fmt.Printf("Action log: %s",dta.somethingA)
}