

func topKFrequent(nums []int, k int) []int {
	fmap := make(map[int]int)
	farr := make([][]int, len(nums))


	for _, num := range nums {
		fmap[num]++
		
		i := fmap[num]

		farr[i-1] = append(farr[i-1], num)
	}


	result := make([]int,0, k)
	u := make(map[int]struct{})
	for i:= len(nums)-1; i>=0; i-- {
		if len(farr[i]) > 0 {
			for _, n := range farr[i] {
				if _, ok := u[n]; !ok {
					result = append(result, n)
					u[n] = struct{}{}
				} 
			}
		}
		if len(result) == k {
				return result
			}
	}

	return result

}
