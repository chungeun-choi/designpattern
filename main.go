package main

import (
	"fmt"

	af "github.com/designpattern/abstractFactory"
	. "github.com/designpattern/builder"
	. "github.com/designpattern/factory"
)


func main() {
	//factory()
	//abstractFactory()
	builder()
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


func abstractFactory() {
	//추상 팩토리인 가구팩토리를 초기화 - Modern으로 생성
	furniture,_ := af.GetFurniture("modern")

	mc := furniture.MakeChair() 
	md := furniture.MakeDesk()
	ms := furniture.MakeSopa()

	fmt.Printf("Moder Chair category: %s",mc.GetCategory())
	fmt.Println()
	fmt.Printf("Moder Desk category: %s",md.GetCategory())	
	fmt.Println()
	fmt.Printf("Moder Sopa category: %s",ms.GetCategory())
}	

func builder() {
	obj := MakeRedisBuilder()

	obj.SetConfig()
	obj.ConnetcObject()
	
	fmt.Print(obj.GetRedisConnection())

	obj.CheckConnection()
}