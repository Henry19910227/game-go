package game

type BeginNewRound struct {
	MiniGameId int
	RoundId    string
	DeckRound  int
	MaxRound   int
}

type BeginDeal struct {
	MiniGameId int
	RoundId    string
	CountDown  int
	Perform    *ActorPerform
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
