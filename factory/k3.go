package factory

type K3 struct {
	Car
}



func newK3() CarProduct {
	return &K3{
		Car: Car{
			name: "K3",
			power: 3,
			maker: "KIA",
		},
	}
}