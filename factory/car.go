package factory

type Car struct {
	name string
	power int
	maker string
}


func (c *Car) setName(name string) {
	c.name = name
}


func(c *Car) setMaker(maker string){
	c.maker = maker
}

func (c *Car) setPower(power int) {
	c.power = power
}

func (c *Car) getName() string{
	return c.name
}

func (c *Car) getPower() int{
	return c.power
}


func (c *Car) getMaker() string {
	return c.maker
}