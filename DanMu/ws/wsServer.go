/*
@Title : wsServer
@Description :
@Author : 谭靖渝
@Update : 2021/8/1 16:35
*/
package ws

// Package ws is to define a websocket server and client connect.
// Author: Arthur Zhang
// Create Date: 20190101

import (
	"DanMu/history"
	"DanMu/pkg/common"
	"encoding/json"
	"github.com/gorilla/websocket"
)

// ws管理器
type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// 客户端
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
	Channel    string
}

// 信息
type Message struct {
	Content   string `json:"content,omitempty"`
}

// 初始化管理器
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

// 开始一个ws服务
func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			info, _ := history.GetInfo(conn.Channel)
			//将历史弹幕发给客户端
			for i := 0; i < len(info.DanMu); i++ {
				jsonMessage, _ := json.Marshal(&Message{Content: info.DanMu[i]})
				conn.Send<-jsonMessage
			}
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			conn.Send<-jsonMessage
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				//取消该redis连接的订阅
				_ = common.S.UnSubscribe(conn.Channel)
				close(conn.Send)
				delete(manager.Clients, conn)
			}
		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}



//读取消息
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Content: string(message)})
		Manager.Broadcast <- jsonMessage
	}
}

//写
func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
