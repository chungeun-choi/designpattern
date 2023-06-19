package abstractfactory



type Desk interface {
	SetName(name string)
	GetName() string
	SetCategory(category string)
	GetCategory() string

}



type EnticDesk struct {
	name string
	category string 
}

func (ec *EnticDesk) SetName(name string) {
	ec.name = name
} 

func (ec *EnticDesk) GetName() string {
	return ec.name
}

func (ec *EnticDesk) SetCategory(category string) {
	ec.category = category
}

func (ec *EnticDesk) GetCategory() string {
	return ec.category
}



type ModernDesk struct {
	name string
	category string 
}

func (ec *ModernDesk) SetName(name string) {
	ec.name = name
} 

func (ec *ModernDesk) GetName() string {
	return ec.name
}

func (ec *ModernDesk) SetCategory(category string) {
	ec.category = category
}

func (ec *ModernDesk) GetCategory() string {
	return ec.category
}