package redis

import (
	"context"
	wsRedis "github.com/sccotJiang/wsj/internal/pkg/redis"
)

var SingleRedisCli *wsRedis.Redis
var ctx = context.Background()

type wsjConf struct {
	SingleRedis struct {
		Ip   string
		Port string
		Db   int
	}
}

func InitRedis() {
	InitwsjRedis()
}

func InitwsjRedis() {
	conf := &wsjConf{SingleRedis: struct {
		Ip   string
		Port string
		Db   int
	}{Ip: "8.142.103.79", Port: "6379", Db: 0}}
	SingleRedisCli = wsRedis.NewRedis(conf.SingleRedis.Port+":"+conf.SingleRedis.Port, conf.SingleRedis.Db)
}
