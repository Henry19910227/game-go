package game

import (
	"game-go/shared/req"
	"game-go/shared/res"
)

type Adapter interface {
	EnterGroup(input *req.EnterGroup) (output *res.GroupInfo, errMsg *res.ErrorMessage)
	EnterGame(input *req.EnterMiniGame) (output *res.EnterMiniGameInfo, errMsg *res.ErrorMessage)
	ClearTrends(input *res.ClearTrends) (output *res.ClearTrends, errMsg *res.ErrorMessage)
	Bet(uid int, input *req.BetReq) (output *req.MyMiniGameBetResult, errMsg *res.ErrorMessage)
	BeginNewRound(input *res.BeginNewRound) (output *res.BeginNewRound, errMsg *res.ErrorMessage)
	BeginDeal(input *res.BeginDeal) (output *res.BeginDeal, errMsg *res.ErrorMessage)
	BeginSettle(input *res.BeginSettle) (output *res.BeginSettle, errMsg *res.ErrorMessage)
}
