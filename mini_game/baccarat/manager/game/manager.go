package game

import (
	gameManager "game-go/shared/mini_game/manager/game"
)

type manager struct {
	gameManager.BaseManager
	pokerMap map[int]int
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

func (m *manager) BetRate(betAreaId int, elements []int) int {
	if len(elements) < 5 {
		return 0
	}
	bankerElements := make([]int, 0)
	tieElements := make([]int, 0)
	var banker int
	var tie int
	index := 0
	// 莊點數計算
	for i, element := range elements {
		if element == -1 {
			index = i // -1 為莊與閒的分界
			break
		}
		point, ok := m.pokerMap[element]
		if !ok {
			return 0
		}
		bankerElements = append(bankerElements, point)
		banker = (banker + point) % 10
	}
	// 閒點數計算
	for i := index + 1; i < len(elements); i++ {
		point, ok := m.pokerMap[elements[i]]
		if !ok {
			return 0
		}
		tieElements = append(tieElements, point)
		tie = (tie + point) % 10
	}
	// 莊
	if betAreaId == 1 {
		if banker > tie {
			return 1
		}
		if banker == tie {
			return 2
		}
		return 0
	}
	// 閒
	if betAreaId == 2 {
		if tie > banker {
			return 1
		}
		if tie == banker {
			return 2
		}
		return 0
	}
	// 和
	if betAreaId == 3 {
		if banker == tie {
			return 1
		}
		return 0
	}
	// 莊對
	if betAreaId == 4 {
		if bankerElements[0] == bankerElements[1] {
			return 1
		}
		return 0
	}
	// 閒對
	if betAreaId == 5 {
		if tieElements[0] == tieElements[1] {
			return 1
		}
		return 0
	}
	// 大
	if betAreaId == 6 {
		if len(bankerElements)+len(bankerElements) > 4 {
			return 1
		}
		return 0
	}
	// 小
	if betAreaId == 7 {
		if len(bankerElements)+len(bankerElements) == 4 {
			return 1
		}
		return 0
	}
	// 幸運六
	if betAreaId == 8 {
		// 莊家以 6 點獲勝，且只用兩張牌 → 賠率 12:1
		if banker > tie && banker == 6 && len(bankerElements) == 2 {
			return 1
		}
		// 莊家以 6 點獲勝，但用了三張牌 → 賠率 20:1
		if banker > tie && banker == 6 && len(bankerElements) == 3 {
			return 2
		}
		return 0
	}
	return 0
}

func (m *manager) WinBetAreaCodes(elements []int) []int {
	winAreaCodes := make([]int, 0)
	for areaCode := 1; areaCode <= 8; areaCode++ {
		if m.BetRate(areaCode, elements) == 0 {
			continue
		}
		winAreaCodes = append(winAreaCodes, areaCode)
	}
	return winAreaCodes
}

func (m *manager) GenerateElements() {
	// 以-1為分界的左半部是莊牌，右半部是閒牌
	m.SetElements([]int{101, 105, 113, -1, 213, 205, 206})
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
