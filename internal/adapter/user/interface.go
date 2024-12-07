package user

import (
	"game-go/internal/model/req"
	"game-go/internal/model/res"
)

type Adapter interface {
	Login(req *req.LoginReq) (*res.InfoAfterLoginSuccess, *res.ErrorMessage)
}
