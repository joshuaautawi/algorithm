package search

// straight forward looping and compare everything
// O(n)
func LinearSearch(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}
