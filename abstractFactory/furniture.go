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
