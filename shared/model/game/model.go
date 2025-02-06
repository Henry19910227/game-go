package game

const (
	Number   int = 1 // 數字
	Poker    int = 2 // 撲克
	Dice     int = 5 // 骰子
	Roulette int = 7 // 輪盤
)

type BeginNewRound struct {
	MiniGameId int
	RoundId    string
	DeckRound  int
	MaxRound   int
}

type BeginDeal struct {
	MiniGameId  int
	RoundId     string
	CountDown   int
	ElementType int
	Performs    []*ActorPerform
}

type ActorPerform struct {
	Elements      []int
	Patterns      []int
	PerformResult []int
}

type BeginSettle struct {
	MiniGameId  int
	RoundId     string
	WinScore    int
	WinAreaCode []int
	Results     []*SettleResult
}

type SettleResult struct {
	AreaCode int
	BetScore int
	WinScore int
}

type SyncAreaBetInfo struct {
	MiniGameId int
	AreaBets   []*AreaBet
}

type AreaBet struct {
	AreaCode  int
	BetScore  int
	UserCount int // 下注的人数
}
