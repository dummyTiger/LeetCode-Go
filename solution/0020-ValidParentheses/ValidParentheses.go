package leetcode

import "LeetCode-Go/utils"

func isValid(s string) bool {

	if len(s) == 0 {
		return true
	}
	stack := utils.NewStack()
	for _, symbol := range s {
		if symbol == 40 || symbol == 123 || symbol == 91 {
			stack.Push(symbol)
		}
		if symbol == 41 {
			value := stack.Pop()
			if value == nil {
				return false
			}
			if v, ok := value.(int32); ok {
				if v != 40 {
					return false
				}
			}
		}

		if symbol == 125 {
			value := stack.Pop()
			if value == nil {
				return false
			}
			if v, ok := value.(int32); ok {
				if v != 123 {
					return false
				}
			}
		}

		if symbol == 93 {
			value := stack.Pop()
			if value == nil {
				return false
			}
			if v, ok := value.(int32); ok {
				if v != 91 {
					return false
				}
			}
		}
	}

	if stack.Len() != 0 {
		return false
	}

	return true

}
