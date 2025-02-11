package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isBaoZi(t *testing.T) {
	elementsArray := [][]int{{1, 1, 1}}
	assert.Equal(t, true, isBaoZi(elementsArray[0]))
	elementsArray = [][]int{{1, 1, 2}}
	assert.Equal(t, false, isBaoZi(elementsArray[0]))
}

func Test_isShunZi(t *testing.T) {
	elementsArray := [][]int{{9, 1, 0}}
	assert.Equal(t, true, isShunZi(elementsArray[0]))
	elementsArray = [][]int{{7, 9, 8}}
	assert.Equal(t, true, isShunZi(elementsArray[0]))
}

func Test_isPair(t *testing.T) {
	elementsArray := [][]int{{5, 5, 0}}
	assert.Equal(t, true, isPair(elementsArray[0]))
	elementsArray = [][]int{{5, 5, 5}}
	assert.Equal(t, false, isPair(elementsArray[0]))
}

func Test_isHalfShun(t *testing.T) {
	elementsArray := [][]int{{9, 1, 0}}
	assert.Equal(t, false, isHalfShun(elementsArray[0]))
	elementsArray = [][]int{{0, 9, 8}}
	assert.Equal(t, false, isHalfShun(elementsArray[0]))
	elementsArray = [][]int{{7, 9, 8}}
	assert.Equal(t, false, isHalfShun(elementsArray[0]))
	elementsArray = [][]int{{9, 2, 0}}
	assert.Equal(t, true, isHalfShun(elementsArray[0]))
}

func Test_isZaLiu(t *testing.T) {
	elementsArray := [][]int{{5, 5, 0}}
	assert.Equal(t, false, isZaLiu(elementsArray[0]))
	elementsArray = [][]int{{0, 5, 9}}
	assert.Equal(t, false, isZaLiu(elementsArray[0]))
	elementsArray = [][]int{{1, 5, 9}}
	assert.Equal(t, true, isZaLiu(elementsArray[0]))
}
