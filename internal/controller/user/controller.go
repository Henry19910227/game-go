package user

import (
	"fmt"
	"game-go/internal/game"
	"game-go/internal/model/req"
	"game-go/internal/model/res"
	"game-go/internal/pkg/tool/crypto"
	"google.golang.org/protobuf/proto"
)

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

func (c *controller) Unmarshal(ctx *game.Context) {
	payload := ctx.MustGet("payload").([]byte)
	var pb = req.LoginReq{}
	if err := proto.Unmarshal(payload, &pb); err != nil {
		ctx.Abort()
		return
	}
	ctx.Set("pb", pb)
}

func (c *controller) Login(ctx *game.Context) {
	loginSuccess := &res.InfoAfterLoginSuccess{
		ResourceBaseUrl: "Henry",
	}
	// 物件轉換成 pb
	pb, err := proto.Marshal(loginSuccess)
	if err != nil {
		fmt.Println(err)
	}
	// 加密
	data, err := crypto.New().Marshal(7, 106, pb)
	if err != nil {
		fmt.Println(err)
	}
	ctx.WriteData(data)
}
