package service

import userService "game-go/internal/service/user"

type Factory interface {
	UserService() userService.Service
}
