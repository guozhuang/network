package service

//定义自己的interface，将整个包的方法规范
type IStartSrv interface {
	Say(string) string
	Speak(string) string
}

type Start struct {
	//Repo repository.IRepo `inject:""`
}

func (s *Start) Say(message string) string {
	return "world"
	//return s.Repo.GetRepoModel().Start.Speak(message)
}

func (s *Start) Speak(message string) string {
	//return s.Repo.GetRepoModel().Start.Speak(message)
	return "world"
}
