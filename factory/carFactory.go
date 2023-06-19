package factory

import "fmt"

func GetCar(carType string) (CarProduct,error){
	if carType =="Genesis" {
		return newGenesis(), nil
	}
	if carType =="K3" {
		return newK3(), nil
	}

	return nil, fmt.Errorf("Wrong car type passed")
}