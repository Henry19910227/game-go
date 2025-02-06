package game

import (
	gameManager "game-go/shared/mini_game/manager/game"
	"sort"
)

type manager struct {
	gameManager.BaseManager
	betMap        map[int]map[int]int
	pokerValueMap map[int]int // 1~10
	pokerPointMap map[int]int // 1~13
}

type NiuType struct {
	MaxSuit  int
	MaxPoint int
	NiuType  int
}

func New(id int, maxRound int) gameManager.Manager {
	m := &manager{}
	m.InitManager(id, maxRound)
	return m
}

func (m *manager) InitManager(id int, maxRound int) {
	m.BaseManager.InitManager(id, maxRound)
	m.InitPokerMap()
	m.InitBetMap()
}

func (m *manager) BetRate(betAreaId int, elementsArray [][]int) int {
	blueNiu := &NiuType{}
	redNiu := &NiuType{}
	for i, elements := range elementsArray {
		if i == 0 {
			// 計算藍方最大牌花色
			blueNiu.MaxSuit = getMaxCardSuit(elements)
			// 計算藍方最大牌點數
			blueNiu.MaxPoint = getMaxCardPoint(elements)
			// 計算牌型
			blueNiu.NiuType = m.GetNiuNiuType(elements)
			continue
		}
		// 計算紅方最大牌花色
		redNiu.MaxSuit = getMaxCardSuit(elements)
		// 計算紅方最大牌花色
		redNiu.MaxPoint = getMaxCardPoint(elements)
		// 計算牌型
		redNiu.NiuType = m.GetNiuNiuType(elements)
	}
	// 藍勝
	if betAreaId == 1 {
		// 比牌型
		if blueNiu.NiuType > redNiu.NiuType {
			return 1
		}
		if blueNiu.NiuType < redNiu.NiuType {
			return 0
		}
		// 平手比最大點數
		if blueNiu.MaxPoint > redNiu.MaxPoint {
			return 1
		}
		if blueNiu.MaxPoint < redNiu.MaxPoint {
			return 0
		}
		// 平手比最大花色(點數越小越好)
		if blueNiu.MaxSuit < redNiu.MaxSuit {
			return 1
		}
		return 0
	}
	// 紅勝
	if betAreaId == 2 {
		// 比牌型
		if blueNiu.NiuType > redNiu.NiuType {
			return 0
		}
		if blueNiu.NiuType < redNiu.NiuType {
			return 1
		}
		// 平手比最大點數
		if blueNiu.MaxPoint > redNiu.MaxPoint {
			return 0
		}
		if blueNiu.MaxPoint < redNiu.MaxPoint {
			return 1
		}
		// 平手比最大花色(點數越小越好)
		if blueNiu.MaxSuit < redNiu.MaxSuit {
			return 0
		}
		return 1
	}
	// 牌型
	return m.betMap[betAreaId][redNiu.NiuType]
}

func (m *manager) WinBetAreaCodes(elementsArray [][]int) []int {
	winAreaCodes := make([]int, 0)
	for areaCode := 1; areaCode <= 26; areaCode++ {
		if m.BetRate(areaCode, elementsArray) == 0 {
			continue
		}
		winAreaCodes = append(winAreaCodes, areaCode)
	}
	return winAreaCodes
}

func (m *manager) GenerateElements() {
	m.SetElements([][]int{{101, 102, 103, 104, 105}, {211, 212, 108, 201, 301}})
}

func (m *manager) InitPokerMap() {
	m.pokerValueMap = make(map[int]int)
	// 黑桃
	getPokerValue(m.pokerValueMap, 101, 113)
	// 紅心
	getPokerValue(m.pokerValueMap, 201, 213)
	// 梅花
	getPokerValue(m.pokerValueMap, 301, 313)
	// 方塊
	getPokerValue(m.pokerValueMap, 401, 413)

	m.pokerPointMap = make(map[int]int)
	// 黑桃
	getPokerPoint(m.pokerPointMap, 101, 113)
	// 紅心
	getPokerPoint(m.pokerPointMap, 201, 213)
	// 梅花
	getPokerPoint(m.pokerPointMap, 301, 313)
	// 方塊
	getPokerPoint(m.pokerPointMap, 401, 413)

}

func (m *manager) InitBetMap() {
	m.betMap = make(map[int]map[int]int)
	// 藍無牛 ~ 藍牛牛
	for i := 3; i <= 13; i++ {
		m.betMap[i] = map[int]int{i - 3: 1} // 0~10
	}
	// 藍特殊牌型
	m.betMap[14] = func() map[int]int { // 11~14
		ma := make(map[int]int)
		for i := 11; i <= 14; i++ {
			ma[i] = 1
		}
		return ma
	}()
	// 紅無牛 ~ 紅牛牛
	for i := 15; i <= 25; i++ {
		m.betMap[i] = map[int]int{i - 15: 1} // 0~10
	}
	// 紅特殊牌型
	m.betMap[26] = func() map[int]int { // 11~14
		ma := make(map[int]int)
		for i := 11; i <= 14; i++ {
			ma[i] = 1
		}
		return ma
	}()
}

func getPokerValue(pokerMap map[int]int, begin int, end int) {
	for i := begin; i <= end; i++ {
		point := i % 100
		if point >= 10 {
			pokerMap[i] = 10
			continue
		}
		pokerMap[i] = point
	}
}

func getPokerPoint(pokerMap map[int]int, begin int, end int) {
	for i := begin; i <= end; i++ {
		point := i % 100
		pokerMap[i] = point
	}
}

// GetSpecialNiuNiuType 計算 四炸牛 (elements = 1~13)
func (m *manager) GetSpecialNiuNiuType(elements []int) int {
	return 0
}

// GetNiuNiuType 計算 牛1~牛9 & 無牛 (elements = 1~10)
func (m *manager) GetNiuNiuType(elements []int) int {
	if len(elements) < 5 {
		return 0
	}
	// 牌面點數 1~13
	pokerPoints := make([]int, 0)
	for _, element := range elements {
		v, ok := m.pokerPointMap[element]
		if !ok {
			return 0
		}
		pokerPoints = append(pokerPoints, v)
	}
	// 牌面價值 1~10
	pokerValues := make([]int, 0)
	for _, element := range elements {
		v, ok := m.pokerValueMap[element]
		if !ok {
			return 0
		}
		pokerValues = append(pokerValues, v)
	}
	// 計算是否是 五小牛
	if ok := isFiveSmallNiu(pokerPoints); ok {
		return 14
	}
	// 計算是否是 五花牛
	if ok := isFiveColorNiu(pokerPoints); ok {
		return 13
	}
	// 計算是否是 四炸牛
	if ok := isFourBlowHiu(pokerPoints); ok {
		return 12
	}
	// 計算是否是 四花牛
	if ok := isFourColorNiu(pokerPoints); ok {
		return 11
	}
	// 計算 無牛~牛牛
	return niuType(pokerValues)
}

func getMaxCardSuit(elements []int) int {
	var maxElement int
	for _, element := range elements {
		if (element % 100) <= (maxElement % 100) {
			continue
		}
		maxElement = element
	}
	return (maxElement / 100) % 10
}

func getMaxCardPoint(elements []int) int {
	var maxPoint int
	for _, element := range elements {
		v := element % 100
		if v <= maxPoint {
			continue
		}
		maxPoint = v
	}
	return maxPoint
}

func isFiveSmallNiu(elements []int) bool {
	var sum int
	for _, element := range elements {
		if element >= 5 {
			return false
		}
		sum += element
	}
	if sum <= 10 {
		return true
	}
	return false
}

func isFiveColorNiu(elements []int) bool {
	for _, element := range elements {
		if element >= 11 {
			continue
		}
		return false
	}
	return true
}
func isFourBlowHiu(elements []int) bool {
	sort.Ints(elements)
	var count int
	var current int
	for _, element := range elements {
		if element != current {
			current = element
			count = 1
			continue
		}
		count++
	}
	return count >= 4
}
func isFourColorNiu(elements []int) bool {
	var count int
	for _, element := range elements {
		if element < 11 {
			continue
		}
		count++
	}
	return count == 4
}

func niuType(elements []int) int {
	sort.Ints(elements)
	cache := make(map[int]int)
	for i, element := range elements {
		cache[i] = element
	}
	for i := 0; i < len(elements)-2; i++ {
		num := elements[i]
		for j := i + 1; j < len(elements)-1; j++ {
			num += elements[j]
			for k := j + 1; k < len(elements); k++ {
				num += elements[k]
				if num%10 != 0 {
					num -= elements[k]
					continue
				}
				delete(cache, i)
				delete(cache, j)
				delete(cache, k)
				var left int
				for _, n := range cache {
					left += n
				}
				ans := left % 10
				if ans == 0 {
					return 10 // 牛牛
				}
				return ans // 1~9 牛
			}
			num -= elements[j]
		}
		num -= elements[i]
	}
	return 0
}
