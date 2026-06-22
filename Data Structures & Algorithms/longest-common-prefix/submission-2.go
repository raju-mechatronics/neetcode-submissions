func longestCommonPrefix(strs []string) string {
    longestPrefix := ""
	i := 0
	j := 0
	for {
		if j >= len(strs[i]) || strs[0][j] != strs[i][j] {
			return longestPrefix
		}
		i=(i+1)%len(strs)
		if i==0 {
			longestPrefix = longestPrefix + string(strs[0][j])
			j++;
		}
	}
	return longestPrefix
}
