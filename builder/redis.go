package builder

import "net/http"







type RedisConnection struct {
	host string
	port int
	auth Auth
	connection *http.Response
	err error
}


func (rc *RedisConnection) SetConfig() {
	rc.host = "loaclhsot.com"
	rc.auth = Auth{
		user: "elastic",
		password: "skscnd331@",
	}
	rc.port = 6379
}

func (rc *RedisConnection) ConnetcObject()  {
	
	resp,err := http.Get(rc.host)

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