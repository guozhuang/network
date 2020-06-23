package model

import (
	"network/model/user"
)

var MCtx = &Model{}

type Model struct {
	UserModel *user.User `auto:"userModel"`
}

func Init() {
	//使用反射进行依赖注入
	/*inject.Register("userModel", &user.User{})

	inject.AutoRegister(MCtx)
	inject.Inject()*/
}
