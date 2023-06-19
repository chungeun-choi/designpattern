package abstractfactory

type Entic interface {
	CreateModernChair(name string, category string)
	CreateModernDesk(name string, category string)
	CreateModernSopa(name string, category string)
}


type EnticProduct struct {
	EnticChair
	EnticDesk
	EnticSopa
}

func (ep *EnticProduct) CreateModernChair(name string, category string) Chair{
	ep.EnticChair = &EnticChair{
		name : name,
		category : categoryl,
	}
	
	return ep.EnticChair
}