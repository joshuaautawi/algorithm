package search

import "fmt"

// the array must already sorted , and basically halving the array to search data
// O(log n)
func BinarySearch(arr []int, n int) bool {
	low := 0
	high := len(arr)
	var mid int
	for {

		mid = low + (high-low)/2
		fmt.Println(mid)
		if n == arr[mid] {
			return true
		} else if n < arr[mid] {
			high = mid
		} else {
			low = mid + 1
		}
		if high < low {
			break
		}
	}
	return false
}
