package arrayFunctions

func StringArrayChunk(strArray []string, size int) [][]string {
	result := [][]string{}

	if 0 == len(strArray) || 0 == size {
		return result
	}

	i := 0
	for _, str := range strArray {
		if i == len(result) {
			result = append(result, []string{})
		}
		result[i] = append(result[i], str)
		if size == len(result[i]) {
			i++
		}
	}

	return result
}

func IntInArray(a []int, x int) bool {
	return Int64InArray(IntArrayToInt64Array(a), int64(x))
}

func Int64InArray(a []int64, x int64) bool {
	if 0 < len(a) {
		for _, i := range a {
			if x == i {
				return true
			}
		}
	}

	return false
}

func Int32InArray(a []int32, x int32) bool {
	return Int64InArray(Int32ArrayToInt64Array(a), int64(x))
}

func IntArrayToInt64Array(a []int) []int64 {
	res := []int64{}

	if 0 < len(a) {
		for _, i := range a {
			res = append(res, int64(i))
		}
	}

	return res
}

func Int32ArrayToInt64Array(a []int32) []int64 {
	res := []int64{}

	if 0 < len(a) {
		for _, i := range a {
			res = append(res, int64(i))
		}
	}

	return res
}
