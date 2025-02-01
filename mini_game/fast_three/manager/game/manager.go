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

func (m *manager) CheckBetResult(betAreaId int, elements []int) int {
	if len(elements) < 10 {
		return 0
	}
	element := elements[0] // 冠軍號碼
	if betAreaId >= 1 && betAreaId <= 21 {
		element = elements[0] + elements[1]   // 冠亞和
		if betAreaId == 3 || betAreaId == 4 { // 冠亞和個位數
			element = element % 10
		}
	}
	return m.betMap[betAreaId][element]
}

func (m *manager) WinBetAreaCodes(elements []int) []int {
	areaCodes := make([]int, 0)
	for betAreaId, areaMap := range m.betMap {
		element := elements[0] // 冠軍號碼
		if betAreaId >= 1 && betAreaId <= 21 {
			element = elements[0] + elements[1] // 冠亞和
			if betAreaId == 3 || betAreaId == 4 {
				element = element % 10 // 冠亞和個位數
			}
		}
		if _, ok := areaMap[element]; !ok {
			continue
		}
		areaCodes = append(areaCodes, betAreaId)
	}
	return areaCodes
}

func (m *manager) GenerateElements() {
	m.SetElements([]int{1, 2, 3})
}

func (m *manager) InitBetMap() {
	m.betMap = make(map[int]map[int]int)
}
