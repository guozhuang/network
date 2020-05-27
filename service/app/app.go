package app

type Service struct {
}

func (s *Service) Say(message string) string {

	return "world"
}
