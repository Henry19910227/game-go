package game

type Manager interface {
	InitManager(id int, maxRound int) // 必要時覆寫

	SetID(id int)
	SetRoundId(roundId string)
	SetDeckRound(deckRoundId int)
	SetMaxRound(maxRound int)
	SetElements(elementsArray [][]int)

	ID() int
	RoundId() string
	DeckRound() int
	MaxRound() int
	NextRound(m Manager) (deckRound int, roundId string)
	Elements() [][]int

	BetRate(betAreaId int, elementsArray [][]int) int // 需要覆寫
	WinBetAreaCodes(elementsArray [][]int) []int      // 需要複寫
	GenerateElements()                                // 需要覆寫
}
