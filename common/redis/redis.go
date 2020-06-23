package redis

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"network/config"
	"time"
)

//还是需要改造成接口的方式，这样onestore的使用同样通过实现接口来实现指定方法的调用
type IRedis interface {
	Get(string) string
}

func NewRedisPool(server string) *redigo.Pool {
	fmt.Println(server)
	return &redigo.Pool{
		MaxIdle:     config.RedisSetting.MaxIdle, //空闲数
		IdleTimeout: config.RedisSetting.IdleTimeout,
		MaxActive:   config.RedisSetting.MaxActive, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server)
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

//当前包进行根据传入的参数初始化，并且提供对外的操作结构，然后进行分别操作。【实现当前的方法接口】

var Commons = &Common{}

type Common struct {
}

//动态绑定当前的结构
func (com *Common) Get(serverAddr string, key string) string {
	fmt.Println(serverAddr)
	conn := NewRedisPool(serverAddr).Get()
	defer conn.Close()

	username, err := redigo.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return username
}
