package service

import (
	"encoding/json"
	"github.com/sccotJiang/wsj/internal/entities/message"
	"github.com/sccotJiang/wsj/internal/entities/namespaces"
	"github.com/sccotJiang/wsj/internal/websocket/entities/client"
	"log"
	"runtime"
	messageService "github.com/sccotJiang/wsj/internal/service/messgae"
)

func sendUserChat(c client.IClient, byteMsg []byte, chatType string) {
	defer func() {
		if err := recover(); err != nil {
			stackSlice := make([]byte, 1024)
			n := runtime.Stack(stackSlice, false)
			log.Printf("sendUserChat %v err:%v,runtime  stack:%s", c.GetId(), err, stackSlice[:n])
		}
	}()
	var chatMsg message.ChatMessage
	err := json.Unmarshal(byteMsg, &chatMsg)
	if err != nil || len(chatMsg.MsgBody) <=0 {
		log.Printf("sendUserChat wrong %v pr MsgBody length %v",err, len(chatMsg.MsgBody))
		return
	}
	if chatType == "publicChat" {//公聊
		messageService.PubMessage(namespaces.CALLER, message.TypeNormal, message.SubTypePublicChat,c.GetId(),byteMsg,[]string{})
	} else {//私聊
		messageService.PubMessage(namespaces.CALLER, message.TypeNormal, message.SubTypePrivateChat,c.GetId(),byteMsg,[]string{chatMsg.TargetUserid})
	}
}
