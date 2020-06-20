package editor

//通过当前结构体进行相应的标准化操作，save和get操作
type User struct {
	UserId   string `user_id`
	UserName string `user_name`
	Phone    string `phone` //经过加密处理之后的信息
	Status   int    `status`
}

func (user *User) Save() {
	//进行editor的save
}
