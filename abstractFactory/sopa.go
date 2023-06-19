package abstractfactory

type SopaInterface interface {
	SetName(name string)
	GetName() string
	SetCategory(category string)
	GetCategory() string

}



type Sopa struct {
	name string
	category string
}

func (s *Sopa) SetName(name string) {
	s.name = name
}

func (s *Sopa) SetCategory(category string) {
	s.category = category
}

func (s *Sopa) GetCategory()string{
	return s.category
}

func (s *Sopa) Getname()string{
	return s.name
}
