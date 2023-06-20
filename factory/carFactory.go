/*
팩토리 메소드 정의는 아래와 같은 순서를 진행합니다

1. 추상 인터페이스를 정의
2. 추상 인터페이스를 바탕으로 구상 클래스를 생성

위와 같은 방법을 통해 확장 가능한 클래스 구조로 생성하여 개발하도록 하는 디자인 패턴입니다
*/

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