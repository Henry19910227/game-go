package game

import (
	gameManager "game-go/shared/mini_game/manager/game"
)

type manager struct {
	gameManager.BaseManager
	pokerMap map[int]int
}

type CardType struct {
	Suit  int
	Point int
}

func New(id int, maxRound int) gameManager.Manager {
	m := &manager{}
	m.InitManager(id, maxRound)
	return m
}

func (m *manager) InitManager(id int, maxRound int) {
	m.BaseManager.InitManager(id, maxRound)
	m.InitPokerMap()
}

func (m *manager) BetRate(betAreaId int, elementsArray [][]int) int {
	longType := &CardType{}
	huType := &CardType{}
	for i, elements := range elementsArray {
		if i == 0 {
			// 計算龍方手牌花色
			longType.Suit = (elements[0] / 100) % 10
			// 計算龍方手牌點數
			longType.Point = elements[0] % 100
			continue
		}
		// 計算虎方手牌花色
		huType.Suit = (elements[0] / 100) % 10
		// 計算虎方手牌點數
		huType.Point = elements[0] % 100
	}
	// 龍勝
	if betAreaId == 1 {
		if longType.Point == huType.Point { // 和局
			return 2
		}
		if longType.Point > huType.Point { // 龍勝
			return 1
		}
		return 0
	}
	// 虎勝
	if betAreaId == 2 {
		if huType.Point == longType.Point { // 和局
			return 2
		}
		if huType.Point > longType.Point { // 虎勝
			return 1
		}
		return 0
	}
	// 和局
	if betAreaId == 3 {
		if huType.Point == longType.Point { // 和局
			return 1
		}
		return 0
	}
	// 龍紅
	if betAreaId == 4 {
		if longType.Suit == 2 || longType.Suit == 4 {
			return 1
		}
		return 0
	}
	// 龍黑
	if betAreaId == 5 {
		if longType.Suit == 1 || longType.Suit == 3 {
			return 1
		}
		return 0
	}
	// 虎紅
	if betAreaId == 6 {
		if huType.Suit == 2 || longType.Suit == 4 {
			return 1
		}
		return 0
	}
	// 虎黑
	if betAreaId == 7 {
		if huType.Suit == 1 || longType.Suit == 3 {
			return 1
		}
		return 0
	}
	return 0
}

func (m *manager) WinBetAreaCodes(elements [][]int) []int {
	winAreaCodes := make([]int, 0)
	for areaCode := 1; areaCode <= 7; areaCode++ {
		if m.BetRate(areaCode, elements) == 0 {
			continue
		}
		winAreaCodes = append(winAreaCodes, areaCode)
	}
	return winAreaCodes
}

func (m *manager) GenerateElements() {
	m.SetElements([][]int{{101}, {201}})
}

func (m *manager) InitPokerMap() {
	m.pokerMap = make(map[int]int)
	// 黑桃
	getPoint(m.pokerMap, 101, 113)
	// 紅心
	getPoint(m.pokerMap, 201, 213)
	// 梅花
	getPoint(m.pokerMap, 301, 313)
	// 方塊
	getPoint(m.pokerMap, 401, 413)
}

func getPoint(pokerMap map[int]int, begin int, end int) {
	for i := begin; i <= end; i++ {
		point := i % 100
		if point >= 10 {
			pokerMap[i] = 0
			continue
		}
		pokerMap[i] = point
	}
}
