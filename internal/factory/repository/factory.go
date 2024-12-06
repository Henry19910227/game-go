package repository

import (
	userRepo "game-go/internal/repository/user"
	"gorm.io/gorm"
)

type factory struct {
	db             *gorm.DB
	userRepository userRepo.Repository
}

func New(db *gorm.DB) Factory {
	repoFactory := &factory{db: db}
	repoFactory.userRepository = userRepo.New(db)
	return repoFactory
}

func (r *factory) UserRepository() userRepo.Repository {
	return r.userRepository
}
