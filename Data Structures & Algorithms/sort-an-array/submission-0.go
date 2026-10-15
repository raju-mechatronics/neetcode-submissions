func QuickSort(arr []int, i, j int) {
	piviot := arr[j]

	l1, l2 := i-1, i

	for l2 < j {
		if arr[l2] <= piviot {
			l1++
			arr[l1], arr[l2] = arr[l2], arr[l1]
		}
		l2++
	}
	arr[l1+1], arr[j] = arr[j], arr[l1+1]
	if i < l1 {
		QuickSort(arr, i, l1)
	}
	if l1+2 < j {
		QuickSort(arr, l1+2, j)
	}
}


func sortArray(nums []int) []int {
	if len(nums) < 2 {return nums}
    QuickSort(nums, 0, len(nums)-1)
	return nums
}
