package abstractfactory


type EnticFurniture struct {
}

type EnticChair struct {
	Chair
}


type EnticDesk struct {
	Desk
}

type EnticSopa struct {
	Sopa
}

func (ef *EnticFurniture) MakeChair() ChairInterface {
	return &EnticChair{
		Chair: Chair{
			name: "entic Chair1",
			category: "entic",
		},
	}
}

func (ef *EnticFurniture) MakeDesk() DeskInterface {
	return &EnticDesk{
		Desk: Desk{
			name: "entic Chair1",
			category: "entic",
		},
	}
}


func (ef *EnticFurniture) MakeSopa() SopaInterface {
	return &EnticSopa{
		Sopa: Sopa{
			name: "entic Chair1",
			category: "entic",
		},
	}
}

