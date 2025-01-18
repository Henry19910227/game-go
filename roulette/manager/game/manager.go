package game

import (
	"math/rand"
	"time"
)

type manager struct {
	id        int
	roundId   string
	deckRound int
	maxRound  int
	elements  []int
	betMap    map[int]map[int]int // 注區id : 對應中獎數字
}

func New(id int, maxRound int) Manager {
	return &manager{id: id, maxRound: maxRound, betMap: createBetMap()}
}

func (m *manager) ID() int {
	return m.id
}

func (m *manager) RoundId() string {
	return m.roundId
}

func (m *manager) DeckRound() int {
	return m.deckRound
}

func (m *manager) MaxRound() int {
	return m.maxRound
}

func (m *manager) NextRound() (deckRound int, roundId string) {
	if m.deckRound == 0 || m.deckRound >= m.maxRound {
		m.deckRound = 0
	}
	m.deckRound++
	m.roundId = time.Now().Format("20060102150405")
	// 設定亂數種子，確保每次執行結果不同
	rand.Seed(time.Now().UnixNano())
	// 生成 0 到 36 之間的隨機數字 (包含 0 和 36)
	//m.elements = []int{rand.Intn(37)}
	m.elements = []int{10}
	return m.deckRound, m.roundId
}

func (m *manager) CheckBetResult(betAreaId int, element int) int {
	if _, ok := m.betMap[betAreaId][element]; !ok {
		return 0
	}
	return 1
}

func (m *manager) WinBetAreaCodes(element int) []int {
	areaCodes := make([]int, 0)
	for index, areaMap := range m.betMap {
		if _, ok := areaMap[element]; !ok {
			continue
		}
		areaCodes = append(areaCodes, index)
	}
	return areaCodes
}

func (m *manager) Elements() []int {
	return m.elements
}

func createBetMap() (betMap map[int]map[int]int) {
	betMap = make(map[int]map[int]int)
	// 小
	smallMap := make(map[int]int)
	for i := 1; i <= 18; i++ {
		smallMap[i] = 1
	}
	// 大
	bigMap := make(map[int]int)
	for i := 19; i <= 36; i++ {
		bigMap[i] = 1
	}
	// 紅
	redMap := make(map[int]int)
	redMap[1] = 1
	redMap[3] = 1
	redMap[5] = 1
	redMap[7] = 1
	redMap[9] = 1
	redMap[12] = 1
	redMap[14] = 1
	redMap[16] = 1
	redMap[18] = 1
	redMap[19] = 1
	redMap[21] = 1
	redMap[23] = 1
	redMap[25] = 1
	redMap[27] = 1
	redMap[30] = 1
	redMap[32] = 1
	redMap[34] = 1
	redMap[36] = 1
	// 黑
	blackMap := make(map[int]int)
	blackMap[2] = 1
	blackMap[4] = 1
	blackMap[6] = 1
	blackMap[8] = 1
	blackMap[10] = 1
	blackMap[11] = 1
	blackMap[13] = 1
	blackMap[15] = 1
	blackMap[17] = 1
	blackMap[20] = 1
	blackMap[22] = 1
	blackMap[24] = 1
	blackMap[26] = 1
	blackMap[28] = 1
	blackMap[29] = 1
	blackMap[31] = 1
	blackMap[33] = 1
	blackMap[35] = 1
	// 單 & 雙
	singleMap := make(map[int]int)
	evenMap := make(map[int]int)
	for i := 1; i <= 36; i++ {
		if i%2 == 0 {
			evenMap[i] = 1
			continue
		}
		singleMap[i] = 1
	}
	// 範圍 1~12
	range1to12Map := make(map[int]int)
	for i := 1; i <= 12; i++ {
		range1to12Map[i] = 1
	}
	// 範圍 13~24
	range13to24Map := make(map[int]int)
	for i := 13; i <= 24; i++ {
		range13to24Map[i] = 1
	}
	// 範圍 25~36
	range25to36Map := make(map[int]int)
	for i := 25; i <= 36; i++ {
		range25to36Map[i] = 1
	}

	betMap[1] = smallMap       // 小
	betMap[2] = bigMap         // 大
	betMap[3] = redMap         // 紅
	betMap[4] = blackMap       // 黑
	betMap[5] = singleMap      // 單
	betMap[6] = evenMap        // 雙
	betMap[7] = range1to12Map  // 範圍 1~12
	betMap[8] = range13to24Map // 範圍 13~24
	betMap[9] = range25to36Map // 範圍 25~36
	// 直注 0 ~ 36
	for i := 10; i <= 46; i++ {
		betMap[i] = map[int]int{i - 10: 1}
	}
	return betMap
}
