package builder

import "net/http"

type Auth struct {
	user string
	password string
	
}

type Conenction struct {
	host string
	auth Auth
	connection *http.Response	
	err error

}