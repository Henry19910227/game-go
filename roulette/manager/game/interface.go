package game

type Manager interface {
	ID() int
	RoundId() string
	DeckRound() int
	MaxRound() int
	NextRound() (deckRound int, roundId string)
	CheckBetResult(betAreaId int, element int) int
	WinBetAreaCodes(element int) []int
	Elements() []int
}
