package game

import (
	"fmt"
	gameAdapter "game-go/core/adapter/game"
	"game-go/core/game"
	"game-go/shared/pkg/tool/crypto"
	"game-go/shared/req"
	"game-go/shared/res"
	"google.golang.org/protobuf/proto"
	"strconv"
)

type controller struct {
	gameAdapter gameAdapter.Adapter
}

func New(gameAdapter gameAdapter.Adapter) Controller {
	return &controller{gameAdapter: gameAdapter}
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
	case mid == 500 && sid == 1005:
		pb = &req.EnterMiniGame{}
	case mid == 500 && sid == 9004:
		pb = &res.BeginNewRound{}
	case mid == 500 && sid == 9010:
		pb = &res.BeginDeal{}
	case mid == 500 && sid == 9005:
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
	// 回傳
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
	// 查看是否有訂閱game，有則全部取消訂閱
	if v, ok := ctx.Client().Get("games"); ok {
		games, _ := v.([]int32)
		for _, gameId := range games {
			ctx.Leave(strconv.Itoa(int(gameId)))
		}
		ctx.Client().Del("games")
	}

	// 離開直播間
	v, ok := ctx.Client().Get("group")
	if !ok {
		return
	}
	groupId, _ := v.(string)
	ctx.Leave(groupId)
	ctx.Client().Del(groupId)

	// 發送響應
	leaveGroup := &res.LeaveGroup{}
	leaveGroup.GroupId = groupId
	pb, _ := proto.Marshal(leaveGroup)
	data, _ := crypto.New().Marshal(500, 1008, pb)
	ctx.WriteData(data)
}

func (c *controller) EnterMiniGame(ctx *game.Context) {
	enterMiniGame := ctx.MustGet("pb").(*req.EnterMiniGame)
	// 加入遊戲
	ctx.Join(strconv.Itoa(int(enterMiniGame.MiniGameId)))
	// 紀錄
	games := make([]int32, 0)
	if v, ok := ctx.Client().Get("games"); ok {
		games = v.([]int32)
	}
	games = append(games, enterMiniGame.MiniGameId)
	ctx.Client().Set("games", games)

	// 回傳
	enterMiniGameInfo := &res.EnterMiniGameInfo{}
	enterMiniGameInfo.MiniGameId = enterMiniGame.MiniGameId
	enterMiniGameInfo.RoundId = "1"
	enterMiniGameInfo.Stage = 1
	enterMiniGameInfo.CountDown = 10
	enterMiniGameInfo.DeckRound = 10

	pb, _ := proto.Marshal(enterMiniGameInfo)
	data, _ := crypto.New().Marshal(500, 1003, pb)
	ctx.WriteData(data)
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
	output, errMsg := c.gameAdapter.BeginNewRound(beginNewRound)
	channelId := strconv.Itoa(int(beginNewRound.MiniGameId))
	if errMsg != nil {
		data, _ := ctx.MarshalData(7, 600, errMsg)
		ctx.Broadcast(channelId, data)
		return
	}
	data, _ := ctx.MarshalData(500, 1004, output)
	ctx.Broadcast(channelId, data)
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
