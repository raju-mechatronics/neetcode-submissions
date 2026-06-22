func getConcatenation(nums []int) []int {
    l := len(nums)
    res := make([]int, 2*l)
    for i:=0; i<(2*l); i++ {
        res[i] = nums[i%l] 
    }
    return res
}
