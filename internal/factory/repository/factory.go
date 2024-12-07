package repository

import (
	userRepo "game-go/internal/repository/user"
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
