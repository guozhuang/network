package controller

import "C"
import "network/controller/user"

var Ctrs = &Ctr{}

//控制器的依赖管理
type Ctr struct {
	User *user.User
}

//应该是被使用的模块来进行初始化挂载
func Init() {
	Ctrs.User = new(user.User)
}
