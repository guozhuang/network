package model

import (
	"network/model/user"
	"network/utils/inject"
)

var Models = &Model{}

type Model struct {
	UserModel *user.User `auto:"userModel"`
}

func Init() {
	//使用反射进行依赖注入
	inject.Register("userModel", &user.User{})

	inject.AutoRegister(Models)
	inject.Inject()
}
