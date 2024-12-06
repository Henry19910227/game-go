package repository

import userRepo "game-go/internal/repository/user"

type Factory interface {
	UserRepository() userRepo.Repository
}
