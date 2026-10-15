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

func Heapify(arr []int, i, j int) {
	c1, c2 := 2*i+1, 2*i+2

	if c1<=j {
		Heapify(arr, c1, j)
	} 

	if c2 <=j {
		Heapify(arr, c2, j)
	}  

	if c1 <= j && arr[i] < arr[c1] {
		arr[c1], arr[i] = arr[i], arr[c1]
	}

	if c2 <= j && arr[i] < arr[c2] {
		arr[c2], arr[i] = arr[i], arr[c2]
	}
}

func HeapSort(arr []int, i, j int) {
	if i<j{
		Heapify(arr, i, j)
		arr[i], arr[j] = arr[j], arr[i]
		HeapSort(arr, i, j-1)
	}
}


func sortArray(nums []int) []int {
	if len(nums) < 2 {return nums}
    HeapSort(nums, 0, len(nums)-1)
	return nums
}
