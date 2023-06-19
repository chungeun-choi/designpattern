package abstractfactory



type DeskInterface interface {
	SetName(name string)
	GetName() string
	SetCategory(category string)
	GetCategory() string

}



type Desk struct {
	name string
	category string
}


func (d *Desk) SetName(name string) {
	d.name = name
}


func (d *Desk) SetCategory(category string) {
	d.category = category
}

func (d *Desk) GetName() string {
	return d.name
}

func (d *Desk) GetCategory() string {
	return d.category 
}