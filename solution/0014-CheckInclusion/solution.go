package _014_CheckInclusion

//todo
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	hash := generateHash(s1)

	exist := true

	for i := 0; i < len(s2)-len(s1)+1; i++ {
		exist = true
		for j := 0; j < len(s1); j++ {
			v := s2[i+j]
			if _, ok := hash[int32(v)]; ok {
				hash[int32(v)] = hash[int32(v)] - 1
			}
		}
		for _, v := range hash {
			if v != 0 {
				exist = false
				break
			}
		}

		if exist {
			return true
		}
		hash = generateHash(s1)
	}

	return exist
}

func generateHash(s1 string) map[int32]int {
	hash := make(map[int32]int)
	for _, v := range s1 {
		if _, ok := hash[v]; ok {
			hash[v] = hash[v] + 1
		} else {
			hash[v] = 1
		}
	}
	return hash
}
