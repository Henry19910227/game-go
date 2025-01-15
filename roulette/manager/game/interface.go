package game

type Manager interface {
	ID() int
	RoundId() string
	DeckRound() int
	MaxRound() int
	NextRound() (deckRound int, roundId string)
	Elements() []int
}
