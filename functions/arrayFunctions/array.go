package arrayFunctions

import "reflect"

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

func InArray(array interface{}, v interface{}) bool {
	if reflect.Slice != reflect.TypeOf(array).Kind() {
		panic("value of 'array' must be a slice")
	}
	a := reflect.ValueOf(array)
	for i := 0; i < a.Len(); i++ {
		if reflect.DeepEqual(v, a.Index(i).Interface()) == true {
			return true
		}
	}

	return false
}

func ArrayChunk(array interface{}, size int) [][]interface{} {
	if reflect.Slice != reflect.TypeOf(array).Kind() {
		panic("value of 'array' must be a slice")
	}
	a := reflect.ValueOf(array)
	result := [][]any{}
	j := 0
	for i := 0; i < a.Len(); i++ {
		if j == len(result) {
			result = append(result, []any{})
		}
		result[j] = append(result[j], a.Index(i).Interface())
		if size == len(result[j]) {
			j++
		}
	}

	return result
}
