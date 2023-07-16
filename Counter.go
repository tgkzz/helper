package helper

func Counter(output []rune) (int, int) {
	lineX := 0
	lineY := 0

	flag := true

	for _, s := range output {
		if s == '\n' {
			lineY++
			flag = false
		} else {
			if flag {
				lineX++
			}
		}
	}

	return lineX, lineY
}
