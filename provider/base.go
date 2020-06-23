package provider

import (
	"fmt"
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
	//FeedRedisProvider *feed.FeedRedisProvider `auto:"feedRedisProvider"`
}

func Init() {
	inject.Register("userRedisProvider", redisFunc(&user.UserRedisProvider{}))
	//inject.Register("feedRedisProvider", redisFunc(&feed.FeedRedisProvider{}))

	//inject.AutoRegister(Ctx)
	inject.Register("provider", Ctx)
	inject.Inject()

	//显然此处并没有实现依赖注入，而是通过全局变量来导出相应的结构来进行的分层调用
	fmt.Println(Ctx.UserRedisProvider.Get("test"))
}

func redisFunc(provider IProvider) IProvider {
	provider.Init()

	return provider
}
