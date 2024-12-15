package user

import (
	"fmt"
	userAdapter "game-go/core/adapter/user"
	"game-go/core/game"
	"game-go/shared/pkg/tool/crypto"
	"game-go/shared/req"
	"google.golang.org/protobuf/proto"
)

type controller struct {
	adapter userAdapter.Adapter
}

func New(adapter userAdapter.Adapter) Controller {
	return &controller{adapter: adapter}
}

func (c *controller) Unmarshal(ctx *game.Context) {
	mid := ctx.MustGet("mid").(uint16)
	sid := ctx.MustGet("sid").(uint16)
	payload := ctx.MustGet("payload").([]byte)
	var pb proto.Message
	switch {
	case mid == 7 && sid == 7:
		pb = &req.LoginReq{}
	default:
		fmt.Println("No matching case for the given mid and sid")
		ctx.Abort()
		return
	}
	if err := proto.Unmarshal(payload, pb); err != nil {
		ctx.Abort()
		return
	}
	ctx.Set("pb", pb)
}

func (c *controller) HeartBeat(ctx *game.Context) {

}

func (c *controller) Login(ctx *game.Context) {
	loginReq := ctx.MustGet("pb").(*req.LoginReq)
	success, errMsg := c.adapter.Login(loginReq)
	if errMsg != nil {
		pb, _ := proto.Marshal(errMsg)
		data, _ := crypto.New().Marshal(7, 107, pb)
		ctx.WriteData(data)
		return
	}
	pb, _ := proto.Marshal(success)
	data, _ := crypto.New().Marshal(7, 106, pb)
	ctx.WriteData(data)
}

func (c *controller) EnterRoom(ctx *game.Context) {
	enterInfo, errMsg := c.adapter.EnterRoom()
	if errMsg != nil {
		pb, _ := proto.Marshal(errMsg)
		data, _ := crypto.New().Marshal(7, 600, pb)
		ctx.WriteData(data)
		return
	}
	pb, _ := proto.Marshal(enterInfo)
	data, _ := crypto.New().Marshal(500, 1000, pb)
	ctx.WriteData(data)
}
