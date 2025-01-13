package repository

import (
	betAreaRepo "game-go/core/repository/bet_area"
	gameRepo "game-go/core/repository/game"
	userRepo "game-go/core/repository/user"
)

type Factory interface {
	UserRepository() userRepo.Repository
	GameRepository() gameRepo.Repository
	BetAreaRepository() betAreaRepo.Repository
}
