package leetcode

import (
	"container/list"
)
//todo wrong
func longestValidParentheses(s string) int {
	stack := list.New()
	count:=0

	max:=0

	for _,ss:=range s {
		if ss == 41 {
			v:=stack.Back()

			if v==nil {
				continue
			}else {
				stack.Remove(v)
				if vv:=v.Value.(int);vv == 40 {
						count = count + 2
						if count>max {
							max = count
						}
				}else {
					count = 0
				}
			}
		} else {
			stack.PushBack(40)
		}
	}

	return max
}
