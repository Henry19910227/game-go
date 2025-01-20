package game

import (
	"fmt"
	gameAdapter "game-go/core/adapter/game"
	"game-go/core/game"
	"game-go/shared/pkg/tool/crypto"
	"game-go/shared/req"
	"game-go/shared/res"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
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
	case mid == 500 && sid == 1030: // 投注 (user 傳入指令)
		pb = &req.BetReq{}
	case mid == 500 && sid == 9011: // 清除歷史紀錄 (game 傳入指令)
		pb = &res.ClearTrends{}
	case mid == 500 && sid == 9004: // 開始新局 (game 傳入指令)
		pb = &res.BeginNewRound{}
	case mid == 500 && sid == 9010: // 開始開牌 (game 傳入指令)
		pb = &res.BeginDeal{}
	case mid == 500 && sid == 9005: // 開始結算 (game 傳入指令)
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
	// API
	output, errMsg := c.gameAdapter.EnterGroup(enterGroup)
	if errMsg != nil {
		data, _ := ctx.MarshalData(7, 600, errMsg)
		ctx.WriteData(data)
		return
	}
	data, _ := ctx.MarshalData(500, 1001, output)
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

	// API
	output, errMsg := c.gameAdapter.EnterGame(enterMiniGame)
	if errMsg != nil {
		data, _ := ctx.MarshalData(7, 600, errMsg)
		ctx.WriteData(data)
		return
	}
	data, _ := ctx.MarshalData(500, 1003, output)
	ctx.WriteData(data)
}

func (c *controller) LeaveMiniGame(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) Bet(ctx *game.Context) {
	betReq := ctx.MustGet("pb").(*req.BetReq)
	uid := ctx.Client().MustGet("uid").(int)

	output, userScore, errMsg := c.gameAdapter.Bet(ctx.MustGet("tx").(*gorm.DB), uid, betReq)
	if errMsg != nil {
		data, _ := ctx.MarshalData(7, 600, errMsg)
		ctx.WriteData(data)
		return
	}
	// 回傳投注結果
	data, _ := ctx.MarshalData(500, 1006, output)
	ctx.WriteData(data)
	// 回傳用戶餘額
	data, _ = ctx.MarshalData(500, 1050, userScore)
	ctx.WriteData(data)
}

func (c *controller) RefreshScore(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) ClearTrends(ctx *game.Context) {
	clearTrends := ctx.MustGet("pb").(*res.ClearTrends)
	output, errMsg := c.gameAdapter.ClearTrends(clearTrends)
	channelId := strconv.Itoa(int(clearTrends.MiniGameIds[0]))
	if errMsg != nil {
		data, _ := ctx.MarshalData(7, 600, errMsg)
		ctx.Broadcast(channelId, data)
		return
	}
	data, _ := ctx.MarshalData(500, 1011, output)
	ctx.Broadcast(channelId, data)
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
	output, errMsg := c.gameAdapter.BeginDeal(beginDeal)
	channelId := strconv.Itoa(int(beginDeal.MiniGameId))
	if errMsg != nil {
		data, _ := ctx.MarshalData(7, 600, errMsg)
		ctx.Broadcast(channelId, data)
		return
	}
	data, _ := ctx.MarshalData(500, 1010, output)
	ctx.Broadcast(channelId, data)
}

func (c *controller) BeginSettle(ctx *game.Context) {
	beginSettle := ctx.MustGet("pb").(*res.BeginSettle)
	channelId := strconv.Itoa(int(beginSettle.MiniGameId))
	clients := ctx.Clients(channelId)
	userIds := make([]int, 0)
	for _, client := range clients {
		value, ok := client.Get("uid")
		if !ok {
			continue
		}
		uid, _ := value.(int)
		userIds = append(userIds, uid)
	}

	settles, errMsg := c.gameAdapter.BeginSettle(beginSettle, userIds)
	if errMsg != nil {
		data, _ := ctx.MarshalData(7, 600, errMsg)
		ctx.Broadcast(channelId, data)
		return
	}

	for _, client := range clients {
		value, ok := client.Get("uid")
		if !ok {
			continue
		}
		uid, _ := value.(int)
		settle, ok := settles[uid]
		if !ok {
			continue
		}
		data, _ := ctx.MarshalData(500, 1005, settle)
		_ = client.Conn().WriteMessage(websocket.BinaryMessage, data)
	}
}
