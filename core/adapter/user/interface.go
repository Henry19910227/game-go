package user

import (
	"game-go/core/model/req"
	"game-go/core/model/res"
)

type Adapter interface {
	Login(req *req.LoginReq) (*res.InfoAfterLoginSuccess, *res.ErrorMessage)
	EnterRoom() (*res.EnterInfo, *res.ErrorMessage)
}
