package user

import (
	"game-go/shared/req"
	"game-go/shared/res"
)

type Adapter interface {
	Login(req *req.LoginReq) (*res.InfoAfterLoginSuccess, *res.ErrorMessage)
	EnterRoom() (*res.EnterInfo, *res.ErrorMessage)
}
