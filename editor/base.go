package editor

import (
	"network/utils/inject"
)

var Ctx = &Editor{}

//包含结构体依旧对应的format方法，之所以隔离该层是因为存在关联结构的符合引用【这样的话，各自的分层model就能包含在各自的包内】
type Editor struct {
	UserEditor UserEditor `auto:"userEditor"`
}

func Init() {
	inject.Register("userEditor", &UserEditor{})

	inject.AutoRegister(Ctx)
	inject.Inject()
}

//获取相应edtor的方法直接共有,返回各自的对应结构
func (editor *Editor) GetUserEditor(userId string) *UserEditor {
	//查询对应数据源，然后初始化相应的editor，并且返回
	user := &UserEditor{
		userId,
		"hah",
		"123344",
		1,
	}

	return user
}

/*func (editor *Editor) GetUserEditors(user_ids []string) []*User {
	//
}*/
