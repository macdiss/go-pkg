package mapFunctions

func StringKeys(v map[string]any) []string {
	result := []string{}
	for k, _ := range v {
		result = append(result, k)
	}

	return result
}
