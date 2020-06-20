package redis

type IRedis interface {
	Get(string) string
	Set(string, string) bool
	MGet([]string) []string
}
