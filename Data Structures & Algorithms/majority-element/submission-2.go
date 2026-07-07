func majorityElement(nums []int) int {
	target:= len(nums)/2
    c := make(map[int]int)
	for _, num := range nums {
		c[num]++
		if c[num] > target {
			return num;
		}
	}
	return nums[0]
}
