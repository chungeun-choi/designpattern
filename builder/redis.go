package builder

import (
	"net/http"
	"strconv"
	"strings"
)







type RedisConnection struct {
	host string
	port int
	auth Auth
	connection *http.Response
	err error
}


func (rc *RedisConnection) SetConfig() {
	rc.host = "localhsot"
	rc.auth = Auth{
		user: "redis",
		password: "skscnd331@",
	}
	rc.port = 6379
}

func (rc *RedisConnection) ConnetcObject()  {
	var stb strings.Builder
	stb.WriteString(rc.host)
	stb.WriteString(":")
	stb.WriteString(strconv.Itoa(rc.port))

	resp,err := http.Get(stb.String())

	rc.connection = resp
	rc.err = err

}

func (rc *RedisConnection) CheckConnection() {

	if rc.err != nil {
        panic(rc.err)
		
    }
}

func MakeRedisBuilder() *RedisConnection{
	return &RedisConnection{}
}


func (rc *RedisConnection) GetRedisConnection() Conenction {

	return Conenction{
		host: rc.host,
		auth: rc.auth,
		connection: rc.connection,
		err: rc.err,
	}
}