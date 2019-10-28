package websocket

import (
	"log"
	"github.com/miRemid/amy/websocket/model"
)

// Client is the struct of cq websocket client
type Client struct{
	baseClient
	message WebEventHandler
}

// WebEventHandler is the handler function of receive message
type WebEventHandler func(event model.CQEvent)

// NewClient return a new websocket server ptr
func NewClient(url string, port int) *Client{
	var client Client
	client.url = url
	client.port = port
	client.token = ""
	return &client
}

// Run will open a websocket client
func (c *Client) Run() {	
	c.Connect("event")
	for {
		_, body, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("read error:%v", err)
		}
		evt := model.NewCQEvent(body)
		go c.message(evt)
	}
}

// OnMessage will set the message parse function
func (c *Client) OnMessage(handler WebEventHandler){
	c.message = handler
}
