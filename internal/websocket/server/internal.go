package server

import (
	"encoding/json"
	"github.com/sccotJiang/wsj/internal/entities/message"
	"github.com/sccotJiang/wsj/internal/entities/namespaces"
	log "github.com/sccotJiang/wsj/internal/pkg/logger"
	"github.com/sccotJiang/wsj/internal/redis"
	"github.com/sccotJiang/wsj/internal/websocket/service"
	"runtime"
)

type InternalManager struct {
	Registers   map[string]chan string //namespace => chan
	UnRegisters map[string]chan string //namespace => chan
}

//InitInternalServer 初始化service相关的组件
func InitInternalServer() {
	instanceIm := &InternalManager{
		Registers:   make(map[string]chan string, len(namespaces.Namespaces)),
		UnRegisters: make(map[string]chan string, len(namespaces.Namespaces)),
	}
	for _, namespace := range namespaces.Namespaces {
		instanceIm.Registers[namespace] = make(chan string, 65530)
		instanceIm.UnRegisters[namespace] = make(chan string, 65530)
		go instanceIm.SubscribeGlobalMsg(namespace)

	}
	service.InitClientManager()
	service.Ini
}
func (im *InternalManager) SubscribeGlobalMsg(namespace string) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("SubscribeGlobalMsg err : %v", err)
		}
	}()
	pubSub := redis.SubscribeGlobalMsg(namespace)
	ch := pubSub.Channel()
	switch namespace {
	case namespaces.CALLER:
		for msg := range ch {
			log.Printf("reveive mmsg:%v", msg.Payload)
			handleMsg(msg.Channel, msg.Payload)
		}
	}
}
func handleMsg(channel string, msg string) {
	defer func() {
		if err := recover(); err != nil {
			stackSlice := make([]byte, 1024)
			s := runtime.Stack(stackSlice, false)
			log.Errorf("handleMsg err %v runtime stack:%s", err, s)
		}
	}()
	var msgList []message.BaseMessage
	if err := json.Unmarshal([]byte(msg), &msgList); err != nil {
		log.Errorf("failed to unmarhsal msg: %v from %v error %v", msg, channel, err)
	}
	for _, msg := range msgList {
		switch msg.Type {
		case message.TypeNormal:
			log.Printf("receive msg:%v", msg.MsgBody)
		}
	}
}
