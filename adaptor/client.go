package adaptor

import "fmt"



type Client struct {
	member Members
}

func (c *Client) Convert(com ConvertIntreaface,path string) {
    c.member = com.ConvertByte(path)
}

func (c *Client) Road(com ConvertIntreaface) {
    fmt.Println("Client inserts Lightning connector into computer.")
    com.RoadObject(&c.member)
}