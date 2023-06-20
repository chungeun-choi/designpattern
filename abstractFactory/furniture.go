/*
디자인 패턴 중 생성과 관련된 내용에서 추상팩토리(abstract factory)에 대한 내용을 실습한 내용입니다
추상 팩토리의 생성 과정은 아래와 같습니다

1. 추상 제품들의 인터페이스 정의
2. 추상 제품 인터페이스 바탕으로 구상 제품 클래스 구현
3. 추상팩토리 인터페이스 정의
4. 개별 추상 팩토리 패턴 클래스 구현
5. 팩토리 메소드를 통한 객체 생성 지원

*/

package abstractfactory

import "fmt"


type FurnitureFactory interface {
	MakeChair() ChairInterface
	MakeDesk() DeskInterface
	MakeSopa() SopaInterface
}

func GetFurniture(style string) (FurnitureFactory,error){
	if style == "entic"{
		return &EnticFurniture{},nil
	} 
	if style == "modern"{
		return &ModernFurniture{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
		
}
