package game

import (
	"math/rand"
	"time"
)

type manager struct {
	id        int
	roundId   string
	deckRound int
	maxRound  int
	elements  []int
}

func New(id int, maxRound int) Manager {
	return &manager{id: id, maxRound: maxRound}
}

func (m *manager) ID() int {
	return m.id
}

func (m *manager) RoundId() string {
	return m.roundId
}

func (m *manager) DeckRound() int {
	return m.deckRound
}

func (m *manager) MaxRound() int {
	return m.maxRound
}

func (m *manager) NextRound() (deckRound int, roundId string) {
	if m.deckRound == 0 || m.deckRound >= m.maxRound {
		m.deckRound = 0
	}
	m.deckRound++
	m.roundId = time.Now().Format("20060102150405")
	// 設定亂數種子，確保每次執行結果不同
	rand.Seed(time.Now().UnixNano())
	// 生成 0 到 36 之間的隨機數字 (包含 0 和 36)
	m.elements = []int{rand.Intn(37)}
	return m.deckRound, m.roundId
}

func (m *manager) Elements() []int {
	return m.elements
}
