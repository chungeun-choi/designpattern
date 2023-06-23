package adaptor

//Member -
type Member struct {
    Name   string
    Age    int
    Active bool
}
 
//Members -
type Members struct {
    Member []Member
}

type ConvertIntreaface interface {
	ConvertByte(path string) *Members
	RoadObject(*Members)
}


