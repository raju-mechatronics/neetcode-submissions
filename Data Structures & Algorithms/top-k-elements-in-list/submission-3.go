

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	fmap := make(map[int]int)
	farr := make([][]int, len(nums))


	for _, num := range nums {
		fmap[num]++
	}


	for k, v:= range fmap {
		farr[v-1] = append(farr[v-1], k)
	}

	result := make([]int,0, k)

	for i:= len(nums)-1; i>=0; i-- {
		if len(farr[i]) > 0 {
			result = append(result, farr[i]...)
			if len(result) == k {
				return result
			}
		}
	}

	return result

}
