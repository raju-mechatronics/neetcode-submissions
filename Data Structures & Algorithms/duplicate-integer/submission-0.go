func hasDuplicate(nums []int) bool {
    numMap := map[int]struct{}{}
    for _, v:=range nums {
        if  _,ok := numMap[v]; ok {
            return true
        }else {
            numMap[v] = struct{}{}
        }
    }
    return false
}
