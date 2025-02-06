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
	m.SetElements([][]int{{101, 102, 103, 104, 105}, {201, 202, 203, 204, 205}})
}

func (m *manager) InitBetMap() {
	m.betMap = make(map[int]map[int]int)
}

func small() map[int]int {
	m := make(map[int]int)
	for i := 3; i <= 10; i++ {
		m[i] = 1
	}
	return m
}

func big() map[int]int {
	m := make(map[int]int)
	for i := 11; i <= 18; i++ {
		m[i] = 1
	}
	return m
}

func single() map[int]int {
	m := make(map[int]int)
	for i := 3; i <= 18; i++ {
		if i%2 == 0 {
			continue
		}
		m[i] = 1
	}
	return m
}
func even() map[int]int {
	m := make(map[int]int)
	for i := 3; i <= 18; i++ {
		if i%2 != 0 {
			continue
		}
		m[i] = 1
	}
	return m
}
func leopard() map[int]int {
	m := make(map[int]int)
	for i := 1; i <= 6; i++ {
		m[i*3] = 1
	}
	return m
}
