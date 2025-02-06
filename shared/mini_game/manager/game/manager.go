package game

import (
	"time"
)

type BaseManager struct {
	id        int
	maxRound  int
	roundId   string
	deckRound int
	elements  [][]int
}

func New(id int, maxRound int) Manager {
	m := &BaseManager{}
	m.InitManager(id, maxRound)
	return m
}

func (b *BaseManager) InitManager(id int, maxRound int) {
	b.id = id
	b.maxRound = maxRound
}

func (b *BaseManager) SetID(id int) {
	b.id = id
}

func (b *BaseManager) SetRoundId(roundId string) {
	b.roundId = roundId
}

func (b *BaseManager) SetDeckRound(deckRoundId int) {
	b.deckRound = deckRoundId
}

func (b *BaseManager) SetMaxRound(maxRound int) {
	b.maxRound = maxRound
}

func (b *BaseManager) SetElements(elements [][]int) {
	b.elements = elements
}

func (b *BaseManager) ID() int {
	return b.id
}

func (b *BaseManager) RoundId() string {
	return b.roundId
}

func (b *BaseManager) DeckRound() int {
	return b.deckRound
}

func (b *BaseManager) MaxRound() int {
	return b.maxRound
}

func (b *BaseManager) NextRound(m Manager) (deckRound int, roundId string) {
	if b.deckRound == 0 || b.deckRound >= b.maxRound {
		b.deckRound = 0
	}
	b.deckRound++
	b.roundId = time.Now().Format("20060102150405")
	m.GenerateElements()
	return b.deckRound, b.roundId
}

func (b *BaseManager) Elements() [][]int {
	return b.elements
}

func (b *BaseManager) BetRate(betAreaId int, elements [][]int) int {
	return 0
}

func (b *BaseManager) WinBetAreaCodes(elements [][]int) []int {
	return []int{}
}

func (b *BaseManager) GenerateElements() {

}
