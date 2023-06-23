package adaptor

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
)




type JsonMembers struct {
}

func MakeJsonObj() *JsonMembers {
	return &JsonMembers{}
}

func (jm *JsonMembers)ConvertByte(path string) *Members {
	fp, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

	data, err := ioutil.ReadAll(fp)
	
	var members *Members

	jsonValue := xml.Unmarshal(data, members)

	if jsonValue != nil {
		panic("Erorr")
	}

	return members

}	



func (jm *JsonMembers)RoadObject(datas *Members){
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
