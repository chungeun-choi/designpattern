package singleton

import (
	"fmt"
	"sync"
)

// type Auth struct {
// 	user string
// 	password string
// }


type Connection struct {
	host string
	port int

}

var lock = &sync.Mutex{}
var ConnValue *Connection

func GetInstance() *Connection {
    if ConnValue == nil {
        lock.Lock()
        defer lock.Unlock()
        if ConnValue == nil {
            fmt.Println("Creating single instance now.")
            
            ConnValue = &Connection{
                host: "192.168.162.1",
                port: 8080,
            }
        } else {
            fmt.Println("Single instance already created.")
        }
    } else {
        fmt.Println("Single instance already created.")
    }

    return ConnValue
}

