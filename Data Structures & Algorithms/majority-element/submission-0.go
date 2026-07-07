func majorityElement(nums []int) int {
    c := make(map[int]int)
	for _, num := range nums {
		c[num]++
	}
	maxval:=0
	result := 0
	for key,value := range c {
		if maxval < value {
			maxval = value
			result = key
		}
	}
	return result
}
