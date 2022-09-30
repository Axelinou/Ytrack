package ToUpper

func ToUpper(text string) string {
	result := ""
	for i := 0; i < len(text); i++ {
		if text[i] > 96 && text[i] < 123 {
			result += string(text[i] - 32)
		} else {
			result += string(text[i])
		}
	}
	return result
}
