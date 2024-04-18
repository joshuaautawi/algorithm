package main

import (
	"algorithm/search"
	"fmt"
)

func main() {
	// arr := []int{1, 2, 3, 4, 5, 6}
	arr1 := []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	res := search.TwoCrystalBall(arr1)
	fmt.Println(res)
}
