package workers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
)

var redisUrl, _ = beego.AppConfig.String("redisurl")

var RedisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.DialURL(redisUrl)
	},
}
