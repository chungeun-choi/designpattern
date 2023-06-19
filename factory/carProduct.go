package factory

type CarProduct interface {
	SetName(name string)
	SetPower(power int)
	SetMaker(maker string)
	GetName() string
	GetPower() int
	GetMaker() string

}