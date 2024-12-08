package repository

import (
	gameRepo "game-go/core/repository/game"
	userRepo "game-go/core/repository/user"
)

type Factory interface {
	UserRepository() userRepo.Repository
	GameRepository() gameRepo.Repository
}
