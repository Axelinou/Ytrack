package IsLower

func IsLower(chaine string) bool {
	for i := 0; i < len(chaine); i++ {
		if chaine[i] < 'a' || chaine[i] > 'z' {
			return false
		}
	}
	return true
}
