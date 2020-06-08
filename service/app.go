package service

type IAppSrv interface {
	Say(string) string
}

type App struct {
}

func (s *App) Say(message string) string {

	return "world!!!"
}
