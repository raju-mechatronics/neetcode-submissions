func twoSum(nums []int, target int) []int {
   tracker := map[int]int{}
   for i, val := range nums {
	if j, ok := tracker[target-val]; ok {
		return []int {j, i}
	} else {
		tracker[val] = i
	}
   } 
   return nil;
}
