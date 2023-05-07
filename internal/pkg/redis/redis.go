package redis

import "github.com/go-redis/redis"

type Redis struct {
	Client *redis.Client
}

func NewRedis(addr string, db int) (r *Redis) {
	r = &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:       addr,
			Password:   "",
			DB:         db,
			MaxRetries: 3,
		}),
	}
	return
}
