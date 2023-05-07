package service

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sccotJiang/wsj/internal/entities/message"
	log "github.com/sccotJiang/wsj/internal/pkg/logger"
	"github.com/sccotJiang/wsj/internal/websocket/entities/client"
	"time"
)

const (
	maxMessageSize = 8192
	pongWait       = 60 * time.Second
)

// ReadClient 读取客户端通过长连接发出来的消息
// PingPong原则：接收客户端的ping，回一个pong，同时更新read和write的deadline
// 接收不到ping会导致read、write channel关闭，从而实现goroutine的回收
func ReadClient(c *client.CjClient) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ReadClient crash %v error %v", c.Id, err)
		}
		GetClientManger().Unregister <- c
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	err := c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return
	}
	c.Conn.SetPingHandler(func(msg string) error {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PingHandler crash %v error %v", c.User.GetUserId(), err)
			}
		}()
		if time.Now().Unix()-c.User.GetLastPing() >= 2 { //如果超过2秒才响应客户端的ping
			c.Conn.SetReadDeadline(time.Now().Add(pongWait))
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
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
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
			sendUserChat(c, msg, baseMsg.Type)
		default:

		}
		time.Sleep(50 * time.Millisecond)
	}
}

func WriteClient(c *client.CjClient) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("WriteClient write crash error:%v,%v", c.Id, err)
		}
	}()
	pingTicker := time.NewTicker(10 * time.Second)
	defer pingTicker.Stop()
	for {
		select {
		case msg, ok := <-c.Sender:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			}
			w, err := c.Conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			c.Conn.SetWriteDeadline(time.Now().Add(pongWait))
			if _, err := w.Write(msg); err != nil {
				log.Errorf("WriteClient %v failed to write msg %s", c.Id, msg)
				return
			}
			if err := w.Close(); err != nil {
				log.Errorf("WriteClient %v write close error:%v", c.Id, err)
				return
			}
		case _, ok := <-c.Ping:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				log.Errorf("WriteClient %v ping chan closed", c.Id)
				return
			}
			if err := c.Conn.WriteMessage(websocket.PongMessage, nil); err != nil {
				log.Errorf("WriteClient %v write pong fail err %v", c.Id, err)
				return
			}
		}
	}
}
