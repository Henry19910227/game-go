package repository

import (
	gameRepo "game-go/internal/repository/game"
	userRepo "game-go/internal/repository/user"
)

type Factory interface {
	UserRepository() userRepo.Repository
	GameRepository() gameRepo.Repository
}
