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

//这里进行初始化的操作并未在相应的依赖加载上生效
func (user *UserRedisProvider) Init() {
	user.ServerAddr = config.RedisAddrSetting.UserRedisAddr
	//fmt.Println("init", user.ServerAddr)
}

func (user *UserRedisProvider) Get(key string) string {
	fmt.Println(user.ServerAddr)
	re := redis.Commons.Get(config.RedisAddrSetting.UserRedisAddr, key)
	return re
}
