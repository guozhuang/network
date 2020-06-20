package user

import (
	"fmt"
	"network/provider"
)

type User struct {
	//
}

func (user *User) OptData(num int) int {
	value := provider.Providers.UserRedis.Get("hello")
	fmt.Println(value)
	return 100 + num
}
