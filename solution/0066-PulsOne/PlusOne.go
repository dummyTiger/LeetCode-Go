package leetcode

func plusOne(digits []int) []int {
	carry := false
	res := make([]int, 0)
	sum := 0
	for i := len(digits) - 1; i >= 0; i = i - 1 {
		if i == len(digits)-1 {
			sum = digits[i] + 1
		} else {
			sum = digits[i]
		}
		if carry {
			sum = sum + 1
		}
		if sum >= 10 {
			sum = sum - 10
			carry = true
		} else {
			carry = false
		}
		res = append(res, sum)
		sum = 0
	}

	if carry {
		res = append(res, 1)
	}

	res1 := make([]int, 0)

	for i := len(res) - 1; i >= 0; i = i - 1 {
		res1 = append(res1, res[i])
	}
	return res1
}
