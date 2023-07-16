package helper

func Atoi(s string) int {
	sign := 1
	res := 0
	dig := 0
	signFound := false
	for i, ch := range s {
		if ch == '+' || ch == '-' {
			if i != 0 || signFound {
				return 0
			}
			if ch == '-' {
				sign = -1
			}
			signFound = true
			continue
		}
		dig = int(ch - '0')
		if dig < 0 || dig > 9 {
			return 0
		}
		res = res*10 + dig
	}
	return res * sign
}
