package builder

import (
	"net/http"
)



type EsConnection struct {
	host string
	port int
	auth Auth
	connection *http.Response
	err error

}



func (ec *EsConnection) SetConfig() {
	ec.host = "loaclhsot.com"
	ec.auth = Auth{
		user: "elastic",
		password: "skscnd331@",
	}
	ec.port = 9200
}

func (ec *EsConnection) ConnetcObject()  {
	
	resp,err := http.Get(ec.host)

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