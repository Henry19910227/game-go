package user

import (
	"game-go/internal/model/req"
	"game-go/internal/model/res"
	"game-go/internal/model/user/login"
	userService "game-go/internal/service/user"
	"strconv"
	"strings"
)

type adapter struct {
	service userService.Service
}

func New(userService userService.Service) Adapter {
	return &adapter{service: userService}
}

func (a *adapter) Login(req *req.LoginReq) (*res.InfoAfterLoginSuccess, *res.ErrorMessage) {
	parts := strings.Split(req.Token, ":")
	errorMessage := &res.ErrorMessage{}
	successMessage := &res.InfoAfterLoginSuccess{}
	if len(parts) < 2 {
		errorMessage.Code = 400
		return nil, errorMessage
	}
	ID, _ := strconv.Atoi(parts[0])
	err := a.service.Login(&login.Input{
		ID:       int64(ID),
		Password: parts[1],
	})
	if err != nil {
		errorMessage.Code = 400
		return nil, errorMessage
	}
	successMessage.ResourceBaseUrl = "Hello world"
	return successMessage, nil
}
