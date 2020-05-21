package start

import (
	"network/repository/start"
)

//声明服务的接口
type IStartService interface {
	Say(message string) string
}

type StartService struct {
	//将要使用的repo进行注入
	Repo start.IStartRepo `inject:""`
}

func (s *StartService) Say(message string) string {
	//这样就实现内部的一层注入写法

	//

	return s.Repo.Speak(message)
}
