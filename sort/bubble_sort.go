package sort

func BubbleSort(arr []int) []int {
	size := len(arr)
	for i := 0; i < size; i++ {
		for y := 0; y < size-1-i; y++ {
			if arr[y] > arr[y+1] {
				temp := arr[y]
				arr[y] = arr[y+1]
				arr[y+1] = temp
			}
		}
	}
	return arr
}
