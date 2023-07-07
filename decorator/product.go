package decorator

import (
	"fmt"
)

type ProductA struct {
	Name string
}

func (pa *ProductA) PrintAnything() {
	
	fmt.Printf("Product Name: %s",pa.Name)
}