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


func (mc *ModernFurniture) MakeChair() ChairInterface {
	return &ModernChair{
		Chair: Chair{
			name: "modern chair1",
			category: "modern",
		},
	}
}


func (mc *ModernFurniture) MakeDesk() DeskInterface {
	return &ModernDesk{
		Desk: Desk{
			name: "modern desk1",
			category: "modern",
		},
	}
}

func (mc *ModernFurniture) MakeSopa() SopaInterface {
	return &ModernSopa{
		Sopa: Sopa{
			name: "modern sopa1",
			category: "modern",
		},
	}
}