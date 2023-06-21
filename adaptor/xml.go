package adaptor

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)


type Memeber struct {
	Name string
	Age int
	Active bool
}


type Members struct {
	Member [] Memeber
}


func ConvertByte() Members {
	fp, err := os.Open("./adaptor/members.xml")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

	data, err := ioutil.ReadAll(fp)

	var members Members

	xmlValue := xml.Unmarshal(data, &members)

	if xmlValue != nil {
		panic("Erorr")
	}

	return members

}	



func RoadObject(datas *Members){
	fp, err := os.Create("./adaptor/convertMembers.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	
}
