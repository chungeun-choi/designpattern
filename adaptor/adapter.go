package adaptor

type Member struct{
	Name string
	Age int
	Active bool	
} 

type Members struct {
	member []Member
}

type ConvertIntreaface interface {
	ConvertByte(path string) Members
	RoadObject(*Members)
}


