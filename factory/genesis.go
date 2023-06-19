package factory



type Genesis struct {
	Car
	
}

func newGenesis() CarProduct {
	return &Genesis{
		Car: Car{
			name: "test",
			power: 3,
			maker: "HYUNDAI",
		},
	}
}