package game

import (
	"game-go/shared/req"
	"game-go/shared/res"
	"gorm.io/gorm"
)

type Adapter interface {
	EnterGroup(input *req.EnterGroup) (output *res.GroupInfo, errMsg *res.ErrorMessage)
	EnterGame(input *req.EnterMiniGame) (output *res.EnterMiniGameInfo, errMsg *res.ErrorMessage)
	ClearTrends(input *res.ClearTrends) (output *res.ClearTrends, errMsg *res.ErrorMessage)
	Bet(tx *gorm.DB, uid int, input *req.BetReq) (output *req.MyMiniGameBetResult, userScore *res.RefreshUserScore, errMsg *res.ErrorMessage)
	BeginNewRound(input *res.BeginNewRound) (output *res.BeginNewRound, errMsg *res.ErrorMessage)
	BeginDeal(input *res.BeginDeal) (output *res.BeginDeal, errMsg *res.ErrorMessage)
	BeginSettle(input *res.BeginSettle, userIds []int) (settles map[int]*res.BeginSettle, errMsg *res.ErrorMessage)
}
