package leetcode

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	position := make(map[int]int)
	res := make([]string, 0)
	for i, j := range score {
		position[j] = i
		res = append(res, strconv.Itoa(j))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(score)))

	for j, i := range score {
		if j == 0 {
			res[position[i]] = "Gold Medal"
		} else if j == 1 {
			res[position[i]] = "Silver Medal"
		} else if j == 2 {
			res[position[i]] = "Bronze Medal"
		} else {
			res[position[i]] = strconv.Itoa(j+1)
		}
	}
	return res

}
