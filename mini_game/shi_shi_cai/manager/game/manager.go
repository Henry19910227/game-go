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
	// 大
	if betAreaId == 1 {
		if sum >= 23 {
			return 1
		}
		return 0
	}
	// 小
	if betAreaId == 2 {
		if sum <= 22 {
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
	// 球1、球2、球3、球4、球5-大
	if betAreaId == 5 || betAreaId == 9 || betAreaId == 13 || betAreaId == 17 || betAreaId == 21 {
		if sum >= 5 {
			return 1
		}
		return 0
	}
	// 球1、球2、球3、球4、球5-小
	if betAreaId == 6 || betAreaId == 10 || betAreaId == 14 || betAreaId == 18 || betAreaId == 22 {
		if sum <= 4 {
			return 1
		}
		return 0
	}
	// 球1、球2、球3、球4、球5-單
	if betAreaId == 7 || betAreaId == 11 || betAreaId == 15 || betAreaId == 19 || betAreaId == 23 {
		if (sum % 2) == 1 {
			return 1
		}
		return 0
	}
	// 球1、球2、球3、球4、球5-雙
	if betAreaId == 8 || betAreaId == 12 || betAreaId == 16 || betAreaId == 20 || betAreaId == 24 {
		if (sum % 2) == 0 {
			return 1
		}
		return 0
	}
	// 定位數字-球1 (在第一球至第五球任意位置上选择1个号码，当该位置的开奖号码与所选号码相同时，即中奖)
	if betAreaId >= 25 && betAreaId <= 34 {
		if _, ok := m.betMap[betAreaId][elementsArray[0][0]]; ok {
			return 1
		}
		return 0
	}
	// 定位數字-球2
	if betAreaId >= 35 && betAreaId <= 44 {
		if _, ok := m.betMap[betAreaId][elementsArray[0][1]]; ok {
			return 1
		}
		return 0
	}
	// 定位數字-球3
	if betAreaId >= 45 && betAreaId <= 54 {
		if _, ok := m.betMap[betAreaId][elementsArray[0][2]]; ok {
			return 1
		}
		return 0
	}
	// 定位數字-球4
	if betAreaId >= 55 && betAreaId <= 64 {
		if _, ok := m.betMap[betAreaId][elementsArray[0][3]]; ok {
			return 1
		}
		return 0
	}
	// 定位數字-球5
	if betAreaId >= 65 && betAreaId <= 74 {
		if _, ok := m.betMap[betAreaId][elementsArray[0][4]]; ok {
			return 1
		}
		return 0
	}
	// 前三-豹子
	if betAreaId == 75 {
		if isBaoZi(elementsArray[0][:3]) {
			return 1
		}
		return 0
	}
	// 前三-順子
	if betAreaId == 76 {
		if isShunZi(elementsArray[0][:3]) {
			return 1
		}
		return 0
	}
	// 前三-對子
	if betAreaId == 77 {
		if isPair(elementsArray[0][:3]) {
			return 1
		}
		return 0
	}
	// 前三-半順
	if betAreaId == 78 {
		if isHalfShun(elementsArray[0][:3]) {
			return 1
		}
		return 0
	}
	// 前三-雜六
	if betAreaId == 79 {
		if isZaLiu(elementsArray[0][:3]) {
			return 1
		}
		return 0
	}
	// 中三-豹子
	if betAreaId == 80 {
		if isBaoZi(elementsArray[0][1:4]) {
			return 1
		}
		return 0
	}
	// 中三-順子
	if betAreaId == 81 {
		if isShunZi(elementsArray[0][1:4]) {
			return 1
		}
		return 0
	}
	// 中三-對子
	if betAreaId == 82 {
		if isPair(elementsArray[0][1:4]) {
			return 1
		}
		return 0
	}
	// 中三-半順
	if betAreaId == 83 {
		if isHalfShun(elementsArray[0][1:4]) {
			return 1
		}
		return 0
	}
	// 中三-雜六
	if betAreaId == 84 {
		if isZaLiu(elementsArray[0][1:4]) {
			return 1
		}
		return 0
	}
	// 後三-豹子
	if betAreaId == 85 {
		if isBaoZi(elementsArray[0][len(elementsArray[0])-3:]) {
			return 1
		}
		return 0
	}
	// 後三-順子
	if betAreaId == 86 {
		if isShunZi(elementsArray[0][len(elementsArray[0])-3:]) {
			return 1
		}
		return 0
	}
	// 後三-對子
	if betAreaId == 87 {
		if isPair(elementsArray[0][len(elementsArray[0])-3:]) {
			return 1
		}
		return 0
	}
	// 後三-半順
	if betAreaId == 88 {
		if isHalfShun(elementsArray[0][len(elementsArray[0])-3:]) {
			return 1
		}
		return 0
	}
	// 後三-雜六
	if betAreaId == 89 {
		if isZaLiu(elementsArray[0][len(elementsArray[0])-3:]) {
			return 1
		}
		return 0
	}
	// 龍
	if betAreaId == 90 {
		if elementsArray[0][0] > elementsArray[0][4] {
			return 1
		}
		return 0
	}
	// 和
	if betAreaId == 91 {
		if elementsArray[0][0] == elementsArray[0][4] {
			return 1
		}
		return 0
	}
	// 虎
	if betAreaId == 92 {
		if elementsArray[0][0] < elementsArray[0][4] {
			return 1
		}
		return 0
	}
	return 0
}

func (m *manager) WinBetAreaCodes(elementsArray [][]int) []int {
	winAreaCodes := make([]int, 0)
	for areaCode := 1; areaCode <= 92; areaCode++ {
		if m.BetRate(areaCode, elementsArray) == 0 {
			continue
		}
		winAreaCodes = append(winAreaCodes, areaCode)
	}
	return winAreaCodes
}

func (m *manager) GenerateElements() {
	m.SetElements([][]int{{1, 2, 3, 4, 5}})
}

func (m *manager) InitBetMap() {
	m.betMap = make(map[int]map[int]int)
	for i := 25; i <= 74; i++ {
		index := (i - 5) % 10
		m.betMap[i] = map[int]int{index: 1}
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
