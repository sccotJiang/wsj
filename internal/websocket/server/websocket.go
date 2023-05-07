package server

import (
	"github.com/gorilla/websocket"
	"github.com/sccotJiang/wsj/internal/entities/namespaces"
	"github.com/sccotJiang/wsj/internal/pkg/logger"
	log "github.com/sccotJiang/wsj/internal/pkg/logger"
	"github.com/sccotJiang/wsj/internal/redis"
	"github.com/sccotJiang/wsj/internal/websocket/entities/client"
	"github.com/sccotJiang/wsj/internal/websocket/service"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//初始化配置
//比如日志的前缀：机器的ip 代码的环境 实例化redis
func InitEnv() {
	var logconfig logger.LogConfig
	logconfig.Ip = "127.0.0.1"
	logger.InitLogger(logconfig)
	log.Println("current ip is: %v", logconfig.Ip)
	redis.InitRedis()
}
func Connnect(manager *service.ClientManager, w http.ResponseWriter, r *http.Request, namespace string) {
	urlParam := r.URL.Query()
	useId := urlParam.Get("uid")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("websocket upgrade errpr:%v", err)
		return
	}
	var c client.CjClient
	switch namespace {
	case namespaces.CALLER:
		c = client.
	}
	manager.Register <- c
}
