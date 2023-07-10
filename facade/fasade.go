package facade

import (
	"fmt"
	"reflect"
)


type FasadeObject struct {
	Ob1 *ObjectOne
	Ob2 *ObjectTwo
	Ob3 *ObjectThreee
	Ob4 *ObjectFour
 }


 func (fo *FasadeObject)Execute() {

	v := reflect.ValueOf(fo)
	elements := v.Elem()
    for i := 0; i < elements.NumField(); i++ {
		mValue := elements.Field(i)
		fmt.Println(mValue.Interface())
		
	}
	
 }
