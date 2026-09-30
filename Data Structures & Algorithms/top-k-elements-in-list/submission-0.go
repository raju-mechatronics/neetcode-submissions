import (
	"maps"
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	ft := make(map[int]int)

	for _, v := range nums {
		ft[v]++
	}

	result := slices.Collect(maps.Keys(ft))

	slices.SortFunc(result, func(a, b int)int {
		return ft[b]-ft[a]
	})
	
	return result[:k]
}
