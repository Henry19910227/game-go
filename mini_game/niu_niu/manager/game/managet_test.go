package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetNiuNiuType(t *testing.T) {
	m := &manager{}
	m.InitManager(10, 50)
	assert.Equal(t, 0, m.GetNiuNiuType([]int{101, 211, 301, 401, 402}))  // 無牛
	assert.Equal(t, 1, m.GetNiuNiuType([]int{309, 402, 313, 102, 308}))  // 牛1
	assert.Equal(t, 5, m.GetNiuNiuType([]int{102, 203, 305, 404, 401}))  // 牛5
	assert.Equal(t, 10, m.GetNiuNiuType([]int{111, 213, 305, 404, 401})) // 牛牛
	assert.Equal(t, 11, m.GetNiuNiuType([]int{111, 212, 410, 311, 413})) // 四花牛
	assert.Equal(t, 12, m.GetNiuNiuType([]int{101, 111, 211, 311, 411})) // 四炸牛
	assert.Equal(t, 13, m.GetNiuNiuType([]int{113, 111, 211, 311, 411})) // 五花牛
	assert.Equal(t, 14, m.GetNiuNiuType([]int{103, 102, 201, 301, 401})) // 五小牛
}

func Test_BetRate(t *testing.T) {
	m := &manager{}
	m.InitManager(10, 50)
	elementsArray := [][]int{{101, 102, 103, 104, 105}, {201, 202, 203, 204, 205}}
	assert.Equal(t, 5, m.GetNiuNiuType(elementsArray[0]))
	assert.Equal(t, 5, m.GetNiuNiuType(elementsArray[1]))
	assert.Equal(t, 1, m.BetRate(1, elementsArray)) // 藍勝
}

func Test_isFourBlowHiu(t *testing.T) {
	elements := []int{111, 211, 401, 311, 411}
	m := &manager{}
	m.InitManager(10, 50)
	// 牌面點數 1~13
	pokerPoints := make([]int, 0)
	for _, element := range elements {
		v, ok := m.pokerPointMap[element]
		if !ok {
			continue
		}
		pokerPoints = append(pokerPoints, v)
	}
	assert.Equal(t, true, isFourBlowHiu(pokerPoints))
}

func Test_GetMaxCardSuit(t *testing.T) {
	assert.Equal(t, 2, getMaxCardSuit([]int{101, 211, 301, 401, 402}))
}
