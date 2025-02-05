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

func (m *manager) BetRate(betAreaId int, elements []int) int {
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
	if len(m.betMap) < 35 {
		return []int{}
	}
	winAreaCodes := make([]int, 0)
	for areaCode := 1; areaCode <= 35; areaCode++ {
		if m.BetRate(areaCode, elements) == 0 {
			continue
		}
		winAreaCodes = append(winAreaCodes, areaCode)
	}
	return winAreaCodes
}

func (m *manager) GenerateElements() {
	m.SetElements([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func (m *manager) InitBetMap() {
	m.betMap = make(map[int]map[int]int)
	// 冠亞和大：12~19
	sumBigMap := make(map[int]int)
	for i := 12; i <= 19; i++ {
		sumBigMap[i] = 1
	}
	// 冠亞和小：3~11
	sumSmallMap := make(map[int]int)
	for i := 3; i <= 11; i++ {
		sumSmallMap[i] = 1
	}
	// 冠亞和單 & 雙：开奖号码中，冠军与亚军相加之和的个位数，其中1、3、5、7、9为“单”，2、4、6、8、0为“双”
	sumSingleMap := make(map[int]int)
	sumEvenMap := make(map[int]int)
	for i := 0; i <= 9; i++ {
		if i%2 == 0 {
			sumEvenMap[i] = 1
			continue
		}
		sumSingleMap[i] = 1
	}
	// 冠軍大：6~10
	bigMap := make(map[int]int)
	for i := 6; i <= 10; i++ {
		bigMap[i] = 1
	}
	// 冠軍小：1~5
	smallMap := make(map[int]int)
	for i := 1; i <= 5; i++ {
		smallMap[i] = 1
	}
	// 冠軍單 & 雙：开奖号码中，冠军号码1、3、5、7、9为“单”，2、4、6、8、10为“双”
	singleMap := make(map[int]int)
	evenMap := make(map[int]int)
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			evenMap[i] = 1
			continue
		}
		singleMap[i] = 1
	}
	// 配置注區
	m.betMap[1] = sumBigMap
	m.betMap[2] = sumSmallMap
	m.betMap[3] = sumSingleMap
	m.betMap[4] = sumEvenMap
	for i := 5; i <= 21; i++ {
		// 冠亞和：选择1个数值（3～19），与开奖号码中冠军、亚军两个号码相加之和一致，即中奖
		m.betMap[i] = map[int]int{i - 2: 1}
	}
	m.betMap[22] = bigMap
	m.betMap[23] = smallMap
	m.betMap[24] = singleMap
	m.betMap[25] = evenMap
	for i := 26; i <= 35; i++ {
		// 冠軍選號：选择一个数字，与开奖号码中冠军号码相同，即中奖
		m.betMap[i] = map[int]int{i - 25: 1}
	}
}
