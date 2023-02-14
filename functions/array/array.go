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
