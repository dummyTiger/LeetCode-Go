package leetcode

import "fmt"

//todo 动态规划 unsolved
func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	dp := make([][]int, 0)

	for i := 0; i < length; i++ {
		arr := make([]int, 0)
		for j := 0; j < length; j++ {
			if j == i {
				arr = append(arr, 1)
			} else {
				arr = append(arr, 0)
			}
		}
		dp = append(dp, arr)
	}

	for i := length - 1; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				n1 := dp[i+1][j]
				n2 := dp[i][j-1]
				if n1 > n2 {
					dp[i][j] = n1
				} else {
					dp[i][j] = n2
				}
			}
			//fmt.Print(fmt.Sprintf("%v",dp[i][j]))
		}
		//fmt.Println()
	}

	size := dp[0][length-1]
	//fmt.Println(size)

	for i,_:=range dp {
		for _,v :=range dp[i] {
			fmt.Print(fmt.Sprintf("%v ",v))
		}
		fmt.Println()
	}





	for i := 0; i <= length-size; i++ {
		temp := s[i : i+size]
		flag := true
		for j, k := 0, len(temp)-1; j <= k; {
			if temp[j] != temp[k] {
				flag = false
				break
			}
			j++
			k--
		}
		if flag {
			return temp
		}
	}
	return ""
}
