package leetcode

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{1})
	result = append(result, []int{1})

	for k := 2; k <= numRows; k++ {
		arr := make([]int, 0)
		temp := result[k-1]
		for i := 0; i < k; i++ {
			if i == 0 {
				arr = append(arr, temp[0])
				continue
			}
			if i == k-1 {
				arr = append(arr, temp[len(temp)-1])
				continue
			}
			arr = append(arr, temp[i-1]+temp[i])
		}
		result = append(result, arr)
	}
	return result[1:]
}
