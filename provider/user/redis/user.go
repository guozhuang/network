package redis

import (
	"fmt"
	redis "network/common/redis"
	"network/config"
)

type IUserRedisProvider interface {
	Get(string) string
}

type UserRedisProvider struct {
	ServerAddr string
}

func (user *UserRedisProvider) Init() {
	user.ServerAddr = config.RedisAddrSetting.UserRedisAddr
	//fmt.Println("init", user.ServerAddr)
	//初始化相应的服务的连接池，然后查看能否进行中转调用【一定可以】
}

func (user *UserRedisProvider) Get(key string) string {
	fmt.Println(user.ServerAddr)
	re := redis.Commons.Get(config.RedisAddrSetting.UserRedisAddr, key)
	return re
}
