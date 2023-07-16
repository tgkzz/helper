package helper

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var isNegative bool
	if n < 0 {
		isNegative = true
		n *= -1
	}

	var result []rune
	for n > 0 {
		digit := n % 10
		n = n / 10
		result = append(result, rune(digit)+'0')
	}

	if isNegative {
		result = append(result, '-')
	}

	// reverse the slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
