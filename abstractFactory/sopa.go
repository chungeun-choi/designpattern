package abstractfactory

type Sopa interface {
	SetName(name string)
	GetName() string
	SetCategory(category string)
	GetCategory() string

}



type EnticSopa struct {
	name string
	category string 
}

func (ec *EnticSopa) SetName(name string) {
	ec.name = name
} 

func (ec *EnticSopa) GetName() string {
	return ec.name
}

func (ec *EnticSopa) SetCategory(category string) {
	ec.category = category
}

func (ec *EnticSopa) GetCategory() string {
	return ec.category
}



type ModernSopa struct {
	name string
	category string 
}

func (ec *ModernSopa) SetName(name string) {
	ec.name = name
} 

func (ec *ModernSopa) GetName() string {
	return ec.name
}

func (ec *ModernSopa) SetCategory(category string) {
	ec.category = category
}

func (ec *ModernSopa) GetCategory() string {
	return ec.category
}