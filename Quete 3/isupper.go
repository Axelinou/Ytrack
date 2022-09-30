package IsUpper

func IsUpper(chaine string) bool {
	for i := 0; i < len(chaine); i++ {
		if chaine[i] < 'A' || chaine[i] > 'Z' {
			return false
		}
	}
	return true
}
