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
	elements := elementsArray[0]
	if len(elements) < 3 {
		return 0
	}
	// 大 小 單 雙 總和
	if betAreaId >= 1 && betAreaId <= 18 {
		// 如果是豹子則通殺
		if !(elements[0] == elements[1] && elements[1] == elements[2]) {
			return m.betMap[betAreaId][elements[0]+elements[1]+elements[2]]
		}
		return 0
	}
	// 單骰
	if betAreaId >= 19 && betAreaId <= 24 {
		var count int
		for _, element := range elements {
			if _, ok := m.betMap[betAreaId][element]; !ok {
				continue
			}
			count++
		}
		return count
	}
	// 對子
	if betAreaId >= 25 && betAreaId <= 30 {
		if elements[0] == elements[1] {
			return m.betMap[betAreaId][elements[0]+elements[1]]
		}
		if elements[1] == elements[2] {
			return m.betMap[betAreaId][elements[1]+elements[2]]
		}
		if elements[0] == elements[2] {
			return m.betMap[betAreaId][elements[0]+elements[2]]
		}
		return 0
	}
	// 豹子點數
	if betAreaId >= 31 && betAreaId <= 36 {
		if elements[0] == elements[1] && elements[1] == elements[2] {
			return m.betMap[betAreaId][elements[0]+elements[1]+elements[2]]
		}
		return 0
	}
	// 豹子
	if betAreaId == 37 {
		if elements[0] == elements[1] && elements[1] == elements[2] {
			return m.betMap[betAreaId][elements[0]+elements[1]+elements[2]]
		}
		return 0
	}
	return 0
}

func (m *manager) WinBetAreaCodes(elementsArray [][]int) []int {
	if len(m.betMap) < 37 {
		return []int{}
	}
	winAreaCodes := make([]int, 0)
	for areaCode := 1; areaCode <= 37; areaCode++ {
		if m.BetRate(areaCode, elementsArray) == 0 {
			continue
		}
		winAreaCodes = append(winAreaCodes, areaCode)
	}
	return winAreaCodes
}

func (m *manager) GenerateElements() {
	m.SetElements([][]int{{6, 6, 5}})
}

func (m *manager) InitBetMap() {
	m.betMap = make(map[int]map[int]int)
	// 大
	m.betMap[1] = big()
	// 小
	m.betMap[2] = small()
	// 單
	m.betMap[3] = single()
	// 雙
	m.betMap[4] = even()
	// 總合
	for i := 5; i <= 18; i++ {
		m.betMap[i] = map[int]int{i - 1: 1} // 總和 4 ~ 17
	}
	// 單骰
	for i := 19; i <= 24; i++ {
		m.betMap[i] = map[int]int{i - 18: 1} // 單骰 1 ~ 6
	}
	// 對子
	m.betMap[25] = map[int]int{2: 1}  // 對子 1
	m.betMap[26] = map[int]int{4: 1}  // 對子 2
	m.betMap[27] = map[int]int{6: 1}  // 對子 3
	m.betMap[28] = map[int]int{8: 1}  // 對子 4
	m.betMap[29] = map[int]int{10: 1} // 對子 5
	m.betMap[30] = map[int]int{12: 1} // 對子 6
	// 豹子
	m.betMap[31] = map[int]int{3: 1}  // 豹子 1
	m.betMap[32] = map[int]int{6: 1}  // 豹子 2
	m.betMap[33] = map[int]int{9: 1}  // 豹子 3
	m.betMap[34] = map[int]int{12: 1} // 豹子 4
	m.betMap[35] = map[int]int{15: 1} // 豹子 5
	m.betMap[36] = map[int]int{18: 1} // 豹子 6
	// 全豹
	m.betMap[37] = leopard()
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
