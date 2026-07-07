func removeElement(nums []int, val int) int {
	cur:= 0
	for i:= 0; i< len(nums); i++ {
		if nums[i] == val {
			continue;
		}
		if i != cur {
			nums[cur] = nums[i]
		}
		cur++
	}
	fmt.Println(cur)
	return cur
}
