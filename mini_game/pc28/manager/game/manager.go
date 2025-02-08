package game

import (
	gameManager "game-go/shared/mini_game/manager/game"
)

type manager struct {
	gameManager.BaseManager
	betMap map[int]map[int]int
}

func New(id int, maxRound int) gameManager.Manager {
	m := &manager{}
	m.InitManager(id, maxRound)
	return m
}

func (m *manager) InitManager(id int, maxRound int) {
	m.BaseManager.InitManager(id, maxRound)
	m.InitBetMap()
}

func (m *manager) BetRate(betAreaId int, elementsArray [][]int) int {
	return 0
}

func (m *manager) WinBetAreaCodes(elementsArray [][]int) []int {
	winAreaCodes := make([]int, 0)
	return winAreaCodes
}

func (m *manager) GenerateElements() {
	m.SetElements([][]int{{6, 6, 6}})
}

func (m *manager) InitBetMap() {
	m.betMap = make(map[int]map[int]int)
}
