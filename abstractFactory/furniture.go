package abstractfactory



type FurnitureFactory interface {
	MakeChair() Chair
	MakeDesk() Desk
	MakeSopa() Sopa
}


// func GetfurnitureFactory(category string) (FurnitureFactory,error){
// 	if category == "entic"{
// 		return &EnticFurniture{}, nill
// 	}
// }