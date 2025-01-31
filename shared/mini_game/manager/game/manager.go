package game

import (
	"math/rand"
	"time"
)

type BaseManager struct {
	id        int
	maxRound  int
	roundId   string
	deckRound int
	elements  []int
	betMap    map[int]map[int]int // 注區id : 對應中獎數字
}

func New(id int, maxRound int) Manager {
	m := &BaseManager{}
	m.InitManager(m, id, maxRound)
	return m
}

func (b *BaseManager) InitManager(m Manager, id int, maxRound int) {
	b.id = id
	b.maxRound = maxRound
	m.InitBetMap()
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

func (b *BaseManager) SetElements(elements []int) {
	b.elements = elements
}

func (b *BaseManager) SetBetMap(betMap map[int]map[int]int) {
	b.betMap = betMap
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

func (b *BaseManager) Elements() []int {
	return b.elements
}

func (b *BaseManager) BetMap() map[int]map[int]int {
	return b.betMap
}

func (b *BaseManager) CheckBetResult(betAreaId int, element int) int {
	if _, ok := b.betMap[betAreaId][element]; !ok {
		return 0
	}
	return 1
}

func (b *BaseManager) WinBetAreaCodes(element int) []int {
	areaCodes := make([]int, 0)
	for index, areaMap := range b.betMap {
		if _, ok := areaMap[element]; !ok {
			continue
		}
		areaCodes = append(areaCodes, index)
	}
	return areaCodes
}

func (b *BaseManager) GenerateElements() {
	// 設定亂數種子，確保每次執行結果不同
	rand.Seed(time.Now().UnixNano())
}

func (b *BaseManager) InitBetMap() {
	b.betMap = make(map[int]map[int]int)
}
