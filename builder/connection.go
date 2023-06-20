package builder

import "net/http"



type Conenction struct {
	host string
	auth Auth
	connection *http.Response	
	err error

}