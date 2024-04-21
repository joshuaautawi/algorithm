package main

import (
	"algorithm/sort"
	"fmt"
)

func main() {
	arr := []int{2, 1, 3, 7, 5, 6}
	// arr1 := []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	res := sort.BubbleSort(arr)
	fmt.Println(res)
}
