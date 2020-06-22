package provider

import (
	redis2 "network/provider/feed/redis"
	"network/provider/user/mongo"
	"network/provider/user/redis"
	"network/utils/inject"
)

var Ctx = &Provider{}

//基础数据层
type Provider struct {
	UserRedisProvider *redis.UserRedisProvider  `auto:"userRedisProvider"`
	UserMongoProvider *mongo.UserMongoProvider  `auto:"userMongoProvider"`
	FeedRedisProvider *redis2.FeedRedisProvider `auto:"feedRedisProvider"`
}

func Init() {
	inject.Register("userRedisProvider", &redis.UserRedisProvider{})
	inject.Register("feedRedisProvider", &redis2.FeedRedisProvider{})

	inject.AutoRegister(Ctx)
	inject.Inject()
}
