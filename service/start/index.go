package start

import (
	"network/repository"
)

//定义自己的interface，将整个包的方法规范
type IStartService interface {
	say(string) string
	Speak(string) string
}

type Service struct {
	Repo repository.Repo
	//Repo repository.IRepo `inject:""`
}

func (s *Service) Say(message string) string {
	return s.Repo.GetRepoModel().Start.Speak(message)
}

func (s *Service) Speak(message string) string {
	return s.Repo.GetRepoModel().Start.Speak(message)
}
