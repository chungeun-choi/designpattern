package adaptor

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)


type XmlMemeber struct {
	Name string
	Age int
	Active bool
}


type XmlMembers struct {
	Member [] XmlMemeber
}

func MakeXmlObj() *XmlMembers {
	return &XmlMembers{}
}

func (members *XmlMembers)ConvertByte(path string) *XmlMembers {
	fp, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

	data, err := ioutil.ReadAll(fp)


	xmlValue := xml.Unmarshal(data, members)

	if xmlValue != nil {
		panic("Erorr")
	}

	return members

}	



func (members *XmlMembers)RoadObject(datas *XmlMembers){
	const path = "./adaptor/convertMembers.xml"
	
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	xmlVluae,xmlErr := xml.Marshal(datas)

	if xmlErr != nil {
		panic(xmlErr)
	}



	ioutil.WriteFile(path,xmlVluae, 0666)

}
