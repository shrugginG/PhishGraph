package utils

func RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, v := range elements {
		if !encountered[v] && v != "" {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}
