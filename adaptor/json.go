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


	jsonValue := json.Unmarshal(data, members)

	if jsonValue != nil {
		panic("Erorr")
	}

	return members

}	



func RoadObject(datas *JsonMembers){
	const path = "./adaptor/convertJson.json"
	
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	jsonValue,jsonErr := json.Marshal(datas)

	if jsonErr != nil {
		panic(jsonErr)
	}



	ioutil.WriteFile(path,jsonValue, 0666)

}
