package redis

//user redis相关操作的代理方法集

type UserRedisProvider struct {
	//
}

//实现对redis的标准化操作以及相应连接池处理

func (user *UserRedisProvider) Get(key string) string {
	return key
}

func (user *UserRedisProvider) Set(key, Value string) string {
	return key
}
