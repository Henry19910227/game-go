package util

func SubtractWithFloor(a int32, b int32) int32 {
	ans := a - b
	if ans < 0 {
		return 0
	}
	return ans
}
