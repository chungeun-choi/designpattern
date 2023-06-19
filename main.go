package main

import (
	"fmt"

	. "github.com/designpattern/factory"
)


func main() {
	factory()
}

func printDetails(c CarProduct){
	fmt.Printf("CAR: %s",c.GetName())
	fmt.Println()
	fmt.Printf("Power: %d",c.GetPower())
	fmt.Println()
	fmt.Printf("Maker: %s",c.GetMaker())
}


func factory() {
	genesis, _ := GetCar("Genesis")
	k3,_ := GetCar("K3")

	printDetails(genesis)
	printDetails(k3)
}