package client

import (
	"github.com/gorilla/websocket"
	"github.com/sccotJiang/wsj/internal/entities/user"
)

//IClient客户端抽象接口
type IClient interface {
	GetId() string
	GetNamespace() string
	GetUser() user.IUser
}
type baseClient struct {
	Id     string
	Conn   *websocket.Conn
	User   user.IUser
	Ping   chan byte
	Sender chan []byte
}

func (c *baseClient) GetId() string {
	return c.Id
}
func (c *baseClient) GetNamespace() string {
	return c.GetNamespace()
}
func (c *baseClient) GetUser() user.IUser {
	return c.User
}
