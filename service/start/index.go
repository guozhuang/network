package start

import (
	"network/repository"
)

var repo repository.Repo

//定义自己的interface，将整个包的方法规范
type IStartService interface {
	say(string) string
	Speak(string) string
}

type Service struct {
	//引入外部的反射接口
	Repo repository.IRepo `inject:""`
}

func (s *Service) Say(message string) string {
	//可以使用链式调用的方式进行处理
	return repo.GetRepoModel().Start.Speak(message)
}

func (s *Service) Speak(message string) string {
	//可以使用链式调用的方式进行处理
	return repo.GetRepoModel().Start.Speak(message)
}
