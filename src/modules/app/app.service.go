package app

type IAppService interface {
	HandlePing() string
	HandleHello() string
}

type AppService struct{}

func NewService() *AppService {
	return &AppService{}
}

func (service *AppService) HandlePing() string {
	return "pong"
}

func (service *AppService) HandleHello() string {
	return "Hello World!"
}
