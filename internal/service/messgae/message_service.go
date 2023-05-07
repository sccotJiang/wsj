package messgae

import (
	"encoding/json"
	"github.com/sccotJiang/wsj/internal/entities/message"
	"github.com/sccotJiang/wsj/internal/redis"
)

//PubToAll pub到消息队列里，让在其他节点上的用户也能收到这次消息
func PubGlobalMessage(namespace, baseType, subType, id string, msg []byte, receiver []string) {

}

// 私信
func PubPrivateMessage(namespace, baseType, subType, id string, msg []byte, receiver string) {
	if len(msg) > 0 {
		tmpMsg := &message.ChatMessage{
			BaseMessage:  message.BaseMessage{Type: baseType},
			TargetUserid: receiver,
		}
		//需要讲消息转换成数组
		jsonMsg, err := json.Marshal([]message.ChatMessage{*tmpMsg})
		if err == nil {
			tmpChannel := redis.GetSubscribeChanIdByUserId(namespace, receiver)
			redis.PublishMsg(namespace, tmpChannel, jsonMsg)
		}
	}
}
