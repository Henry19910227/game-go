package repository

import (
	gameRepo "game-go/core/repository/game"
	userRepo "game-go/core/repository/user"
	"gorm.io/gorm"
)

type factory struct {
	db *gorm.DB
}

func New(db *gorm.DB) Factory {
	repoFactory := &factory{db: db}
	return repoFactory
}

func (r *factory) UserRepository() userRepo.Repository {
	return userRepo.New(r.db)
}

func (r *factory) GameRepository() gameRepo.Repository {
	return gameRepo.New(r.db)
}
