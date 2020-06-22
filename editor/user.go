package editor

import "fmt"

//通过当前结构体进行相应的标准化操作，save和get操作
type UserEditor struct {
	UserId   string `user_id`
	UserName string `user_name`
	Phone    string `phone` //经过加密处理之后的信息
	Status   int    `status`
}

func (user *UserEditor) Save() {
	//进行editor的save:将字段转换成标准的形式，转换的方法标准化处理
	fmt.Println("save", user.UserId)
}
