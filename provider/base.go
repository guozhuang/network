package provider

import (
	feed "network/provider/feed/redis"
	user "network/provider/user/redis"
	"network/utils/inject"
)

type IProvider interface {
	Init()
}

var Ctx = &Provider{}

//基础数据层
type Provider struct {
	UserRedisProvider *user.UserRedisProvider `auto:"userRedisProvider"`
	//UserMongoProvider *mongo.UserMongoProvider  `auto:"userMongoProvider"`
	FeedRedisProvider *feed.FeedRedisProvider `auto:"feedRedisProvider"`
}

func (provider *Provider) Init() {
	//
}

func Init() {
	inject.Register("userRedisProvider", redisFunc(&user.UserRedisProvider{}))
	inject.Register("feedRedisProvider", redisFunc(&feed.FeedRedisProvider{}))

	inject.Register("provider", Ctx)
	inject.Inject()

	//fmt.Println(Ctx.UserRedisProvider.Get("test"))
}

func redisFunc(provider IProvider) IProvider {
	provider.Init()

	return provider
}
