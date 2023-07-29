package helper

func ToUpper(s string) string {
	res := ""

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			res += string(ch - 'a' + 'Z')
		} else {
			res += string(ch)
		}
	}

	return res
}
