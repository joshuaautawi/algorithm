package search

import (
	"math"
)

// given two crystal ball that will break if drop from high enough floor
// determine the exact spot in which it will break
// example arr := []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}

// O sqrt(n)
func TwoCrystalBall(arr []int) int {
	size := len(arr)
	jmpAmount := int(math.Floor(math.Sqrt(float64(size))))
	var startIndex int
	// first crystal the break range
	for i := jmpAmount; i < size; i += jmpAmount {
		if arr[i] == 1 {
			startIndex = i - int(jmpAmount)
			break
		}
	}
	for i := 0; i < jmpAmount && i < size; i++ {
		if arr[startIndex] == 1 {
			return startIndex
		}
		startIndex++
	}
	return -1

}
