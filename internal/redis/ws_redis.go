package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/sccotJiang/wsj/internal/entities/namespaces"
	"strconv"
)

var subscribePreKey = "subscribe_pre"

func SubscribeGlobalMsg(namespace string) *redis.PubSub {
	var pubSub *redis.PubSub
	switch namespace {
	case namespaces.CALLER:
		SingleRedisCli.Client.Subscribe()
	}
	return pubSub
}

func GetSubscribeChanIdByUserId(namespace, userId string) string {
	switch namespace {
	case namespaces.CALLER:
		num, _ := strconv.Atoi(userId)
		return strconv.Itoa(num % 20)
	default:
		return ""
	}
}
func PublishMsg(namespace, channelId string, msg []byte) {
	switch namespace {
	case namespaces.CALLER:
		SingleRedisCli.Client.Publish(subscribePreKey+channelId, msg)
	}
}
