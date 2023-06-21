package builder

import (
	"net/http"
	"strconv"
	"strings"
)



type EsConnection struct {
	host string
	port int
	auth Auth
	connection *http.Response
	err error

}



func (ec *EsConnection) SetConfig() {
	ec.host = "127.0.0.1"
	ec.auth = Auth{
		user: "elastic",
		password: "skscnd331@",
	}
	ec.port = 9200
}

func (ec *EsConnection) ConnetcObject()  {
	var stb strings.Builder
	stb.WriteString(ec.host)
	stb.WriteString(":")
	stb.WriteString(strconv.Itoa(ec.port))

	resp,err := http.Get(stb.String())

	ec.connection = resp
	ec.err = err

}

func (ec *EsConnection) CheckConnection() {

	if ec.err != nil {
        panic(ec.err)
		
    }
}

func MakeEsBuilder() *EsConnection{
	return &EsConnection{}
}


func (ec *EsConnection) GetEsConnection() Conenction {

	return Conenction{
		host: ec.host,
		auth: ec.auth,
		connection: ec.connection,
		err: ec.err,
	}
}