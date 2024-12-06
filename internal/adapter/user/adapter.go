package user

import "game-go/internal/service/user"

type adapter struct {
	service user.Service
}

func New(userService user.Service) Adapter {
	return &adapter{service: userService}
}
