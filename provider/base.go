package provider

import (
	"fmt"
	redis2 "network/provider/feed/redis"
	"network/provider/user/mongo"
	"network/provider/user/redis"
	"network/utils/inject"
)

var Providers = &Provider{}

//基础数据层
type Provider struct {
	UserRedis *redis.UserRedisProvider  `auto:"userRedis"`
	UserMongo *mongo.UserMongoProvider  `auto:"userMongo"`
	FeedRedis *redis2.FeedRedisProvider `auto:"feedRedis"`
}

func Init() {
	//始终还是需要进行注册之后进行依赖注入
	inject.Register("userRedis", &redis.UserRedisProvider{})
	inject.Register("feedRedis", &redis2.FeedRedisProvider{})

	inject.AutoRegister(Providers)
	inject.Inject()

	//Providers.User.Redis.Get("hello")
	fmt.Println(Providers.UserRedis.Get("hi"))

}

/*func ProviderFunc(provider IProvider) *Provider {
	//进行标准的类型转换，最后返回对应的providers总结构也是一个接口类型即可
	return provider.(*Provider)
}

//自动绑定相应元素的mongo和redis的字段
func NewProvider(provideName string, provider IProvider) IProvider {
	//查找当前结构体中存在mongo或是存在redis字段，并且进行实例化绑定
	//查找当前结构内的字段
	elems := reflect.ValueOf(provider)
	fmt.Println(elems)
	t := elems.Type()

	fmt.Println("type:", t.Name())

	for i := 0; i < t.NumField(); i ++ {
		field := t.Field(i)
		value := elems.Field(i).Interface()

		switch valueType := value.(type) {
		case int:
			fmt.Printf("%s: %v = %d\n", field.Name, field.Type, valueType)//
		case string:
			fmt.Printf("%s: %v = %s\n", field.Name, field.Type, valueType)
		default:
			fmt.Printf("%s: %v = %s\n", field.Name, field.Type, valueType)
		}
	}

	/*switch provideName {
	case "user":
		tmpProvider := new(user.Provider)
		tmpProvider.Redis = new(redis.UserRedisProvider)
		provider.User = tmpProvider
		return provider
	default:
		return provider
	}
	return provider
}*/
