package model

import "network/model/user"

var Models = &Model{}

type Model struct {
	User *user.User
}

func Init() {
	//todo:后续考虑将此处的初始化动态化处理
	Models.User = new(user.User)
}
