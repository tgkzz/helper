package helper

func CharFrequency(s string) map[string]int {
	res := make(map[string]int)

	for _, ch := range s {
		res[string(ch)]++
	}

	return res
}
