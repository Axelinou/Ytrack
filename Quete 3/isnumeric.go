package isnumeric

func IsNumeric(text string) bool {
	for i := 0; i < len(text); i++ {
		if !(text[i] > 47 && text[i] < 58) {
			return false
		}
	}
	return true
}
