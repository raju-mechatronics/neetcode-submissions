import "slices"

func topKFrequent(nums []int, k int) []int {
	fmap := make(map[int]int)
	farr := make([][]int, len(nums))


	for _, num := range nums {
		fmap[num]++
		
		i := fmap[num]

		farr[i-1] = append(farr[i-1], num)

		if i-2 >= 0 {
			index := slices.Index(farr[i-2], num)
			farr[i-2] = slices.Delete(farr[i-2], index, index+1)
		}
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
