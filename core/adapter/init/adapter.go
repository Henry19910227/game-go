package init

import initService "game-go/core/service/init"

type adapter struct {
	service initService.Service
}

func New(service initService.Service) Adapter {
	return &adapter{service: service}
}

func (a *adapter) InitBetAreaCache() (err error) {
	return a.service.InitBetAreaCache()
}
