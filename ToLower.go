package helper

func ToLower(s string) string {
	res := ""

	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			res += string(ch - 'A' + 'a')
		} else {
			res += string(ch)
		}
	}

	return res
}
