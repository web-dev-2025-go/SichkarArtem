package service

type Service struct {
	HelloService HelloService
}

func NewService(HelloService HelloService) *Service {
	return &Service{
		HelloService: HelloService,
	}
}
