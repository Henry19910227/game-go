package game

import (
	gameManager "game-go/shared/mini_game/manager/game"
)

type manager struct {
	gameManager.BaseManager
}

func New(id int, maxRound int) gameManager.Manager {
	m := &manager{}
	m.InitManager(m, id, maxRound)
	return m
}

func (m *manager) CheckBetResult(betAreaId int, element int) int {
	return 0
}

func (m *manager) WinBetAreaCodes(element int) []int {
	return []int{}
}

func (m *manager) GenerateElements() {
	m.BaseManager.GenerateElements()
	m.SetElements([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func (m *manager) InitBetMap() {
	m.BaseManager.InitBetMap()
}
