package tools

//item in slice?
func InSlice(items []int64, item int64) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
func DeleteSlice(a []int64, elem int64) []int64 {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

func removeRepByLoop(slc []int64) []int64 {
	var result []int64
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}

func removeRepByMap(slc []int64) []int64 {
	var result []int64
	tempMap := map[int64]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func RemoveRep(slc []int64) []int64 {
	if len(slc) < 1024 {
		return removeRepByLoop(slc)
	} else {
		return removeRepByMap(slc)
	}
}
