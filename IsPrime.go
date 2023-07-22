package helper

func Binarray(a []int) int {
	max := 0

	if len(a) == 1 {
		return 0
	}

	if len(a) == 2 {
		if a[0] != a[1] {
			return 2
		}
	}

	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			count := 1
			for j := i + 1; j < len(a); j++ {
				if a[j] != a[j-1] {
					count++
				} else {
					break
				}
			}
			if count > max {
				max = count
			}
		}
	}

	return max
}
