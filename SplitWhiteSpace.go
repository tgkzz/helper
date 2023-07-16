package helper

func SplitWhiteSpaces(s string) []string {
	res := []string{}
	wordCollector := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\n' && s[i] != '\t' {
			wordCollector += string(s[i])
		} else if wordCollector != "" {
			res = append(res, wordCollector)
			wordCollector = ""
		}
	}
	if wordCollector != "" {
		res = append(res, wordCollector)
	}
	return res
}
