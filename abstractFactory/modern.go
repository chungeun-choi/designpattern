package abstractfactory


type ModernChair struct {
	Chair
}

type ModernDesk struct {
	Desk
}
type ModernSopa struct {
	Sopa
}

type ModernFurniture struct {}


func (mc *ModernFurniture) MakeChair() Chair {
	return &ModernChair{
		Chair: Chair{
			name: "modern chair",
			category: "",
		},
	}
}