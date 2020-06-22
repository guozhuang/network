package model

import (
	"network/model/user"
	"network/utils/inject"
)

var Ctx = &Model{}

type Model struct {
	UserModel *user.User `auto:"userModel"`
}

func Init() {
	//使用反射进行依赖注入
	inject.Register("userModel", &user.User{})

	inject.AutoRegister(Ctx)
	inject.Inject()
}
