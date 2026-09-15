import (
    "slices"
    "maps"
)

// space o(1) time o(n)
func MakeCharFrequency(word string) [26]int {
    freqTab := [26]int{}
    for _, r := range word {
        freqTab[r-'a']++
    }
    return freqTab
}


func groupAnagrams(strs []string) [][]string {
    freqMap := make(map[[26]int][]string)

    for _, word := range strs {
        curfrqTbl := MakeCharFrequency(word)
        freqMap[curfrqTbl] = append(freqMap[curfrqTbl], word)
    }

    return slices.Collect(maps.Values(freqMap))
}
