package facade

import (
	"fmt"
)


type PackageInterface interface {
	PrintObject()
}



type ObjectOne struct {
	Name string
}

func (oo *ObjectOne)PrintObject() {
	fmt.Println(oo.Name)
}


type ObjectTwo struct {
	Name string
}

func (ot *ObjectTwo)PrintObject() {
	fmt.Println(ot.Name)
}


type ObjectThreee struct {
	Name string
}

func (ot *ObjectThreee)PrintObject() {
	fmt.Println(ot.Name)
}

type ObjectFour struct {
	Name string
}

func (of *ObjectFour)PrintObject() {
	fmt.Println(of.Name)
}