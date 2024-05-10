package math

// get sum from  range a to b
func GetSum(a, b int) int {
	if a > b {
		a, b = b, a
	}
	return (a + b) * (b - a + 1) / 2
}
