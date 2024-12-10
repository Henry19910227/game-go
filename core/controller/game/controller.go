package game

import (
	"fmt"
	"game-go/core/game"
	"game-go/core/model/req"
	"game-go/core/model/res"
	"game-go/core/pkg/tool/crypto"
	"google.golang.org/protobuf/proto"
	"strconv"
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
	var pb proto.Message
	switch {
	case mid == 500 && sid == 1001:
		pb = &req.EnterGroup{}
	case mid == 500 && sid == 1003:
		ctx.Next()
		return
	case mid == 500 && sid == 1004:
		pb = &res.BeginNewRound{}
	case mid == 500 && sid == 1010:
		pb = &res.BeginDeal{}
	case mid == 500 && sid == 1005:
		pb = &res.BeginSettle{}
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

func (c *controller) EnterGroup(ctx *game.Context) {
	enterGroup := ctx.MustGet("pb").(*req.EnterGroup)
	// 加入直播間
	ctx.Join(enterGroup.IdP)
	// 紀錄
	ctx.Client().Set("group", enterGroup.IdP)

	gameInfo := &res.MiniGameBasicInfo{}
	gameInfo.MiniGameId = 9
	gameInfo.CountDown = 20
	gameInfo.Stage = 1
	gameInfo.Name = "輪盤"

	groupInfo := &res.GroupInfo{}
	groupInfo.MiniGameBasicInfoList = []*res.MiniGameBasicInfo{gameInfo}
	pb, _ := proto.Marshal(groupInfo)
	data, _ := crypto.New().Marshal(500, 1001, pb)
	ctx.WriteData(data)
}

func (c *controller) LeaveGroup(ctx *game.Context) {
	// 取出紀錄
	v, ok := ctx.Client().Get("group")
	if !ok {
		return
	}
	groupId, ok := v.(string)
	if !ok {
		return
	}
	// 離開直播間
	ctx.Leave(groupId)
	// 刪除紀錄
	ctx.Client().Del(groupId)

	leaveGroup := &res.LeaveGroup{}
	leaveGroup.GroupId = groupId
	pb, _ := proto.Marshal(leaveGroup)
	data, _ := crypto.New().Marshal(500, 1008, pb)
	ctx.WriteData(data)
}

func (c *controller) EnterMiniGame(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) LeaveMiniGame(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) Bet(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) RefreshScore(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) BeginNewRound(ctx *game.Context) {
	beginNewRound := ctx.MustGet("pb").(*res.BeginNewRound)
	beginNewRound.RoundId = "202412091500"
	beginNewRound.DeckRound = 1
	gameId := strconv.Itoa(int(beginNewRound.MiniGameId))

	pb, _ := proto.Marshal(beginNewRound)
	data, _ := crypto.New().Marshal(500, 1004, pb)
	ctx.Broadcast(gameId, data)
}

func (c *controller) BeginDeal(ctx *game.Context) {
	beginDeal := ctx.MustGet("pb").(*res.BeginDeal)
	beginDeal.RoundId = "202412091500"
	gameId := strconv.Itoa(int(beginDeal.MiniGameId))

	pb, _ := proto.Marshal(beginDeal)
	data, _ := crypto.New().Marshal(500, 1010, pb)
	ctx.Broadcast(gameId, data)
}

func (c *controller) BeginSettle(ctx *game.Context) {
	beginSettle := ctx.MustGet("pb").(*res.BeginSettle)
	beginSettle.WinAreaCodes = []int32{1, 2, 3}
	beginSettle.WinScore = 1000
	gameId := strconv.Itoa(int(beginSettle.MiniGameId))

	pb, _ := proto.Marshal(beginSettle)
	data, _ := crypto.New().Marshal(500, 1005, pb)
	ctx.Broadcast(gameId, data)
}
