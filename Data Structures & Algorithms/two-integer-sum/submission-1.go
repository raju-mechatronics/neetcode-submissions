func twoSum(nums []int, target int) []int {
    memo := make(map[int]int)

	for i, num := range nums {
		n2 := target - num
		if j, ok := memo[n2]; ok {
			return []int{j, i}
		} else {
			memo[num] = i
		}
	}
	return nil
}
