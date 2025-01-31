package game

type Manager interface {
	InitManager(m Manager, id int, maxRound int)

	SetID(id int)
	SetRoundId(roundId string)
	SetDeckRound(deckRoundId int)
	SetMaxRound(maxRound int)
	SetElements(elements []int)
	SetBetMap(betMap map[int]map[int]int)

	ID() int
	RoundId() string
	DeckRound() int
	MaxRound() int
	NextRound(m Manager) (deckRound int, roundId string)
	Elements() []int
	BetMap() map[int]map[int]int

	CheckBetResult(betAreaId int, element int) int // 需要覆寫
	WinBetAreaCodes(element int) []int             // 需要複寫
	GenerateElements()                             // 需要覆寫
	InitBetMap()                                   // 需要覆寫
}
