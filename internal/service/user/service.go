package user

import "game-go/internal/repository/user"

type service struct {
	repository user.Repository
}

func New(userRepository user.Repository) Service {
	return &service{repository: userRepository}
}
