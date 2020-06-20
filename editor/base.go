package editor

//包含结构体依旧对应的format方法，之所以隔离该层是因为存在关联结构的符合引用【这样的话，各自的分层model就能包含在各自的包内】
type Editor struct {
	User
}

//获取相应edtor的方法直接共有,返回各自的对应结构
/*func (editor *Editor) GetUserEditor(user_id string) *User {
	//
}

func (editor *Editor) GetUserEditors(user_ids []string) []*User {
	//
}*/
