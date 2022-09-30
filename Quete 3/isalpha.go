package isalpha

func IsAlpha(text string) bool {
	for i := 0; i < len(text); i++ {
		if !((text[i] < 123 && text[i] > 64 && (text[i] > 96 || text[i] < 91)) || (text[i] > 47 && text[i] < 58)) {
			return false
		}
	}
	return true
}
