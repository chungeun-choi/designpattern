package abstractfactory



type ChairInterface interface {
	SetName(name string)
	GetName() string
	SetCategory(category string)
	GetCategory() string

}



type Chair struct {
	name string
	category string
}


func (c *Chair) SetName(name string){
	c.name = name
}


func (c *Chair) SetCategory(category string){
	c.category = category
} 

func (c *Chair) GetName() string{
	return c.name
}

func (c *Chair) GetCategory() string{
	return c.category
}