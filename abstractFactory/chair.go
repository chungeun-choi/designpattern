package abstractfactory



type Chair interface {
	SetName(name string)
	GetName() string
	SetCategory(category string)
	GetCategory() string

}



type EnticChair struct {
	name string
	category string 
}

func (ec *EnticChair) SetName(name string) {
	ec.name = name
} 

func (ec *EnticChair) GetName() string {
	return ec.name
}

func (ec *EnticChair) SetCategory(category string) {
	ec.category = category
}

func (ec *EnticChair) GetCategory() string {
	return ec.category
}



type ModernChair struct {
	name string
	category string 
}

func (ec *ModernChair) SetName(name string) {
	ec.name = name
} 

func (ec *ModernChair) GetName() string {
	return ec.name
}

func (ec *ModernChair) SetCategory(category string) {
	ec.category = category
}

func (ec *ModernChair) GetCategory() string {
	return ec.category
}