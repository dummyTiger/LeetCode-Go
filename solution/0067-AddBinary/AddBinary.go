package leetcode

import "strings"

func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)

	res := make([]string, 0)
	carry := false
	i := 0

	for {
		idxA := lenA - i - 1
		idxB := lenB - i - 1

		if idxA >= 0 && idxB >= 0 {
			valueA := a[idxA]
			valueB := b[idxB]

			if valueA == 48 && valueB == 48 {
				if !carry {
					res = append(res, "0")
					carry = false
					i++
				} else {
					res = append(res, "1")
					carry = false
					i++
				}
			} else if valueA == 49 && valueB == 49 {
				if !carry {
					res = append(res, "0")
					carry = true
					i++
				} else {
					res = append(res, "1")
					carry = true
					i++
				}
			} else if valueA == 48 && valueB == 49 {
				if !carry {
					res = append(res, "1")
					carry = false
					i++
				} else {
					res = append(res, "0")
					carry = true
					i++
				}
			} else if valueA == 49 && valueB == 48 {
				if !carry {
					res = append(res, "1")
					carry = false
					i++
				} else {
					res = append(res, "0")
					carry = true
					i++
				}
			}
			continue
		}

		if idxA < 0 && idxB < 0 {
			break
		}

		if idxA < 0 && idxB >= 0 {
			if b[idxB] == 48 {
				if carry {
					res = append(res, "1")
					carry = false
				} else {
					res = append(res, "0")
					carry = false
				}
			} else if b[idxB] == 49 {
				if carry {
					res = append(res, "0")
					carry = true
				} else {
					res = append(res, "1")
					carry = false
				}
			}
			i++
			continue
		}

		if idxB < 0 && idxA >= 0 {
			if a[idxA] == 48 {
				if carry {
					res = append(res, "1")
					carry = false
				} else {
					res = append(res, "0")
					carry = false
				}
			} else if a[idxA] == 49 {
				if carry {
					res = append(res, "0")
					carry = true
				} else {
					res = append(res, "1")
					carry = false
				}
			}
			i++
			continue
		}

	}

	if carry {
		res = append(res, "1")
	}

	res1 := make([]string, 0)

	for i := len(res) - 1; i >= 0; i = i - 1 {
		res1 = append(res1, res[i])
	}

	return strings.Join(res1, "")
}
