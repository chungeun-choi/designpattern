package factory

type Car struct {
	name string
	power int
	maker string
}


func (c *Car) SetName(name string) {
	c.name = name
}


func(c *Car) SetMaker(maker string){
	c.maker = maker
}

func (c *Car) SetPower(power int) {
	c.power = power
}

func (c *Car) GetName() string{
	return c.name
}

func (c *Car) GetPower() int{
	return c.power
}


func (c *Car) GetMaker() string {
	return c.maker
}