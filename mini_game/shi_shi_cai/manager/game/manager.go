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
	return []int{1}
}

func (m *manager) GenerateElements() {
	m.SetElements([][]int{{1, 2, 3, 4, 5}})
}

func (m *manager) InitBetMap() {

}
