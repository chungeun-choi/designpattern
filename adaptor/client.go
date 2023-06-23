package adaptor

import "fmt"



type Client struct {
	data *Members
}

func (c *Client) Convert(com ConvertIntreaface,path string) {
    c.data = com.ConvertByte(path)
    fmt.Println(c.data)
}

func (c *Client) Road(com ConvertIntreaface) {
    fmt.Println("Client inserts Lightning connector into computer.")
    com.RoadObject(c.data)
}