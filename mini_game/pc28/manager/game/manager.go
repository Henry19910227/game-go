package game

import (
	gameManager "game-go/shared/mini_game/manager/game"
	"sort"
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
	var sum int
	for _, element := range elementsArray[0] {
		sum += element
	}
	// 豹子
	if betAreaId == 33 {
		if isBaoZi(elementsArray[0]) {
			return 1
		}
		return 0
	}
	// 豹子通殺
	if isBaoZi(elementsArray[0]) {
		return 0
	}
	// 大
	if betAreaId == 1 {
		if sum >= 14 {
			return 1
		}
		return 0
	}
	// 小
	if betAreaId == 2 {
		if sum <= 13 {
			return 1
		}
		return 0
	}
	// 單
	if betAreaId == 3 {
		if (sum % 2) == 1 {
			return 1
		}
		return 0
	}
	// 雙
	if betAreaId == 4 {
		if (sum % 2) == 0 {
			return 1
		}
		return 0
	}
	// 和值
	if betAreaId >= 5 && betAreaId <= 32 {
		if _, ok := m.betMap[betAreaId][sum]; ok {
			return 1
		}
		return 0
	}
	// 順子
	if betAreaId == 34 {
		if isShunZi(elementsArray[0]) {
			return 1
		}
		return 0
	}
	// 對子
	if betAreaId == 35 {
		if isPair(elementsArray[0]) {
			return 1
		}
		return 0
	}
	// 半順
	if betAreaId == 36 {
		if isHalfShun(elementsArray[0]) {
			return 1
		}
		return 0
	}
	// 雜六
	if betAreaId == 37 {
		if isZaLiu(elementsArray[0]) {
			return 1
		}
		return 0
	}
	return 0
}

func (m *manager) WinBetAreaCodes(elementsArray [][]int) []int {
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
	m.SetElements([][]int{{6, 6, 6}})
}

func (m *manager) InitBetMap() {
	m.betMap = make(map[int]map[int]int)
	for i := 5; i <= 32; i++ {
		m.betMap[i] = map[int]int{i - 5: 1}
	}
}

// 判斷豹子 (3个相同的号码)
func isBaoZi(elements []int) bool {
	current := elements[0]
	for _, element := range elements {
		if current == element {
			continue
		}
		return false
	}
	return true
}

// 判斷順子 (3个相连的号码，顺序不限)
func isShunZi(elements []int) bool {
	sort.Ints(elements) // 先排序，確保順序
	if elements[0] == 0 && elements[1] == 8 && elements[2] == 9 {
		return true
	}
	if elements[0] == 0 && elements[1] == 1 && elements[2] == 9 {
		return true
	}
	if elements[2]-elements[0] == 2 {
		return true
	}
	return false
}

// 判斷對子 (3个数字中有1个数字出现两次，另1一个数字出现1次，顺序不限)
func isPair(elements []int) bool {
	sort.Ints(elements)
	var count int
	for i := 0; i < len(elements)-1; i++ {
		if elements[i+1]-elements[i] > 0 {
			continue
		}
		count++
	}
	return count == 1
}

// 判斷半順 (3个数字各不相同，且其中只有两个数字相连，顺序不限)
func isHalfShun(elements []int) bool {
	sort.Ints(elements)
	// 判斷三個元素是否都不相同
	for i := 0; i < len(elements)-1; i++ {
		if elements[i+1]-elements[i] > 0 {
			continue
		}
		return false
	}
	// 判斷頭尾是否連續
	if elements[0] == 0 && elements[2] == 9 {
		// 判斷中間元素是否與頭尾連續
		if elements[1]-elements[0] == 1 || elements[2]-elements[1] == 1 {
			return false
		}
		return true
	}
	// 判斷其他兩個元素連續
	var count int
	for i := 0; i < len(elements)-1; i++ {
		if elements[i+1]-elements[i] != 1 {
			continue
		}
		count++
	}
	return count == 1
}

// 判斷雜六 (3个数字各不相同且互不相连)
func isZaLiu(elements []int) bool {
	sort.Ints(elements)
	// 判斷頭尾是否連續
	if elements[0] == 0 && elements[2] == 9 {
		return false
	}
	// 判斷其他數字是否連續
	for i := 0; i < len(elements)-1; i++ {
		if elements[i+1]-elements[i] > 1 {
			continue
		}
		return false
	}
	return true
}
