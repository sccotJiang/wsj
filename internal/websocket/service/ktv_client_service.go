package service

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sccotJiang/wsj/internal/entities/message"
	"github.com/sccotJiang/wsj/internal/websocket/entities/client"
	"log"
	"time"
)

const (
	maxMessageSize = 8192
	pangWait       = 60 * time.Second
)

//读取客户端通过长连接发出来的消息
func ReadClient(c *client.CjClient) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ReadClient crash %v error %v", c.Id, err)
		}
		GetClientManger().Unregister <- c
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	err := c.Conn.SetReadDeadline(time.Now().Add(pangWait))
	if err != nil {
		return
	}
	c.Conn.SetPingHandler(func(msg string) error {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PingHandler crash %v error %v", c.User.GetUserId(), err)
			}
		}()
		if time.Now().Unix()-c.User.GetLastPing() >= 2 {
			c.Conn.SetReadDeadline(time.Now().Add(pangWait))
			c.User.SetLastPing(time.Now().Unix())
			c.Ping <- 'a'
		}
		return nil
	})
	c.Conn.SetPongHandler(func(msg string) error {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PongHandler crash %v error %v", c.User.GetUserId(), err)
			}
		}()
		c.Conn.SetReadDeadline(time.Now().Add(pangWait))
		c.User.SetLastPing(time.Now().Unix())
		return nil
	})
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				log.Printf("ReadClient normal closure %v error %v", c.Id, err)
				break
			}
			continue
		}
		var baseMsg message.BaseMessage
		err = json.Unmarshal(msg, baseMsg)
		if err != nil {
			log.Printf("ReadClient parse client %v message %s error…%v", c.Id, msg, err)
		}
		switch baseMsg.Type {
		case "publicchat":
			
		default:

		}
	}
}
