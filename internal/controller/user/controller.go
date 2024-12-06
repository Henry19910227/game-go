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

func New() Controller {
	return &controller{}
}

func (c *controller) Unmarshal(ctx *game.Context) {
	mid := ctx.MustGet("mid").(uint16)
	sid := ctx.MustGet("sid").(uint16)
	payload := ctx.MustGet("payload").([]byte)
	if mid == 7 && sid == 7 {
		var pb = req.LoginReq{}
		if err := proto.Unmarshal(payload, &pb); err != nil {
			ctx.Abort()
			return
		}
		ctx.Set("pb", &pb)
	}
}

func (c *controller) HeartBeat(ctx *game.Context) {

}

func (c *controller) Login(ctx *game.Context) {
	loginReq := ctx.MustGet("pb").(*req.LoginReq)
	loginSuccess := &res.InfoAfterLoginSuccess{
		ResourceBaseUrl: loginReq.Nickname,
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

func (c *controller) EnterRoom(ctx *game.Context) {
	enterRoom := &res.EnterInfo{
		ProtoVersion: "1.0",
		GameConfigs: []*res.GameConfig{
			{
				MiniGameId: 1,
				BetAreaConfigs: []*res.BetAreaConfig{
					{
						AreaCode: 1,
						Name:     "",
						Odds:     []float32{100},
						MinLimit: 1000,
						MaxLimit: 1000000,
					},
				},
			},
		},
	}
	// 物件轉換成 pb
	pb, err := proto.Marshal(enterRoom)
	if err != nil {
		fmt.Println(err)
	}
	// 加密
	data, err := crypto.New().Marshal(500, 1000, pb)
	if err != nil {
		fmt.Println(err)
	}
	ctx.WriteData(data)
}
