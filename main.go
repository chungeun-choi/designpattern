package main

import (
	"fmt"

	af "github.com/designpattern/abstractFactory"
	. "github.com/designpattern/adaptor"
	. "github.com/designpattern/builder"
	. "github.com/designpattern/factory"

	//. "github.com/designpattern/iterator"
	. "github.com/designpattern/decorator"
	. "github.com/designpattern/facade"
	. "github.com/designpattern/observer"
	. "github.com/designpattern/prototype"
	. "github.com/designpattern/singleton"
	. "github.com/designpattern/strategy"
)


func main() {
	fmt.Println("Testing")
	//factory()
	//abstractFactory()
	//builder()
	//adaptor()
	//observer()
	//strategy()
	//prototype()
	//singleton()
	//decorator()
	facade()
	
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
	obj := MakeEsBuilder()

	obj.SetConfig()
	obj.ConnetcObject()
	
	fmt.Print(obj.GetEsConnection())

	obj.CheckConnection()
}


func adaptor() {
	xmlObj := MakeXmlObj()
	jsonObj := MakeJsonObj()

	client := &Client{}

	client.Convert(xmlObj,"/Users/cucuridas/Desktop/study_project/designpattern/adaptor/members.xml")
	client.Road(jsonObj)
	// fmt.Println(client.member)

}



func observer(){
	store := NewStore("Nike")
	customer1 := &Customer{
		ProductExists: false,
		Id:"이제되네",
	}

	customer2 := &Customer{
		ProductExists: false,
		Id: "이제되네2",
	}
	store.Register(customer1)
	store.Register(customer2)

	store.UpdateProdcut()
}


func strategy() {
	
	obj := &Numbser{
		Parameter:  []int{1, 2, 3},
	}
	strategyObj := &Plus{}
	obj.SetStrategy(strategyObj)
	result := obj.DoSomething()

	fmt.Println(result)

}


func prototype() {
	colorRobot := &ColorRobot{
		Color: "RED",
	}
	armRobot := &ArmRobot{
		Children: colorRobot,
		Arm: "Is Settings",
	}


	armRobot.Running()
}


func singleton() {
	for i := 0; i < 30; i++ {
        go GetInstance()
    }

	fmt.Scanln()
}

func decorator() {
	product := &ProductA{
		Name: "Choi",
	}
	decoAObj := &DecoratorTypeA{
		Product: *product,
	}
	decoBObj := &DecoratorTypeB{
		Product: *product,
	}

	decoAObj.GetProduct()
	decoBObj.GetProduct()
}

func facade() {
	fo := &FasadeObject{
		Ob1: &ObjectOne{
			Name: "test1",
		},
		Ob2: &ObjectTwo{
			Name: "test1",
		},
		Ob3: &ObjectThreee{
			Name: "test1",
		},
		Ob4: &ObjectFour{
			Name: "test1",
		},
	}

	fo.Execute()
}