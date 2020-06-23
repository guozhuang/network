package redis

import (
	"network/config"
)

//user redis相关操作的代理方法集

type FeedRedisProvider struct {
	ServerAddr string
}

func (feed *FeedRedisProvider) Init() {
	feed.ServerAddr = config.RedisSetting.Host
}

func (feed *FeedRedisProvider) Get(key string) string {
	return key
}
