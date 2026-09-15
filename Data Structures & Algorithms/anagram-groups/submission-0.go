func MakeCharFrequency(word string) []int {
    freqTab := make([]int, 26, 26)
    for _, r := range word {
        freqTab[r-'a']++
    }
    return freqTab
}

func Match(l, r []int) bool {
    for i:= range 26 {
        if l[i] != r[i] {
            return false
        }
    }
    return true
}

func groupAnagrams(strs []string) [][]string {
    frequencyGroup := make([][]int, 0)

    result := make([][]string, 0)

    for _, word := range strs {
        curfrqTbl := MakeCharFrequency(word)
      
        found := false
        for i, frqTbl := range frequencyGroup {
            if Match(curfrqTbl, frqTbl) {
                result[i] = append(result[i], word)
                found = true
                break
            }
        } 
        if !found {
            frequencyGroup= append(frequencyGroup, curfrqTbl)
            result = append(result, []string{word})
        }
    }

    return result
}
