package redis

//user redis相关操作的代理方法集

type UserRedisProvider struct {
	//
}

func (user *UserRedisProvider) Get(key string) string {
	return key
}

func (user *UserRedisProvider) Set(key, Value string) string {
	return key
}
