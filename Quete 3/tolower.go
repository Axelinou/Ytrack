package tolower

func ToLower(text string) string {
	result := ""
	for i := 0; i < len(text); i++ {
		if text[i] > 64 && text[i] < 91 {
			result += string(text[i] + 32)
		} else {
			result += string(text[i])
		}
	}
	return result
}
