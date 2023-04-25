package service

import (
	"github.com/sccotJiang/wsj/internal/entities/namespaces"
	"github.com/sccotJiang/wsj/internal/websocket/entities/client"
	"log"
	"sync"
)

type ClientManager struct {
	clientLists map[string]*sync.Map
	Register    chan client.IClient //建立连接
	Unregister  chan client.IClient //断开连接
}

var instanceCM *ClientManager

func GetClientManger() *ClientManager {
	return instanceCM
}

func (cm *ClientManager) ManageClient() {

	for {
		select {
		case c := <-cm.Register: //连接上来的用户
			namespace := c.GetNamespace()
			cm.clientLists[namespace].Store(c.GetId(), c)
			log.Printf("manager Register for %v:%v", namespace, c.GetId())
			switch namespace {
			case namespaces.CALLER:
			}
			//case c := <-cm.Unregister: //断开连接的用户
		}
	}
}

func InitClientManager() {
	instanceCM = &ClientManager{
		clientLists: make(map[string]*sync.Map, len(namespaces.Namespaces)),
		Register:    make(chan client.IClient, 63500),
		Unregister:  make(chan client.IClient, 63500),
	}
	go instanceCM.ManageClient()
}

func GraceTerminate(namespace string, wg *sync.WaitGroup) {
	defer wg.Done()
}
