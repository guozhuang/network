package start

import (
	"fmt"
)

//用来测试使用反射的依赖注入实现

//定义一个接口
type IStartRepo interface {
	Speak(message string) string
}

//实现接口
type StartRepo struct {
	//Source datasource.IDb `inject:""`//数据源连接
}

//当前简单实现直接返回，不使用数据源
func (s *StartRepo) Speak(message string) string {
	return fmt.Sprintf("[Repository] speak: %s", message)
}
