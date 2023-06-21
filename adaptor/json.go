package adaptor

import (
	"encoding/json"
	"io/ioutil"
	"os"
)


type JsonMemeber struct {
	Name string
	Age int
	Active bool
}


type JsonMembers struct {
	Member [] JsonMembers
}

func MakeJsonObj() *JsonMembers {
	return &JsonMembers{}
}

func (members *JsonMembers)ConvertByte(path string) *JsonMembers {
	fp, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

	data, err := ioutil.ReadAll(fp)


	xmlValue := json.Unmarshal(data, members)

	if xmlValue != nil {
		panic("Erorr")
	}

	return members

}	



func RoadObject(datas *JsonMembers){
	const path = "./adaptor/convertMembers.xml"
	
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	xmlVluae,xmlErr := json.Marshal(datas)

	if xmlErr != nil {
		panic(xmlErr)
	}



	ioutil.WriteFile(path,xmlVluae, 0666)

}
