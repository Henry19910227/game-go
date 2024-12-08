package service

import userService "game-go/core/service/user"

type Factory interface {
	UserService() userService.Service
}
