package service

import (
	"sync"
)

type ClientManager struct {
	clientLists map[string]*sync.Map
}

var instanceCM *ClientManager

func InitClientManager() {
	instanceCM = &ClientManager{
		clientLists: make(map[string]*sync.Map, len(namespace)),
	}
}
