package composite

import (
	"fmt"
)

type Node struct {
	nodes []ComponentInterface
	name string
}


func (n *Node)Search(){
	fmt.Printf("Node name is %s",n.name)
	for _,v := range(n.nodes){
		v.Search()
	}
}