package redis

//user redis相关操作的代理方法集

type FeedRedisProvider struct {
	//
}

func (feed *FeedRedisProvider) Get(key string) string {
	return key
}

func (feed *FeedRedisProvider) Set(key, Value string) string {
	return key
}
