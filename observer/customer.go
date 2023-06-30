package observer

import "fmt"


type Customer struct {
	Id string
	ProductExists bool
}



func (c *Customer)Update(product string){
	fmt.Println("상품이 들어왔습니다!")
	c.ProductExists = true
}

func (c *Customer)GetID() string {
	return c.Id
}

