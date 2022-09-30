package Join

func Join(tab []string, a string) string {
	phrase := ""
	for i := 0; i < len(tab); i++ {
		phrase = phrase + tab[i]
		if i != len(tab)-1 {
			phrase = phrase + a
		}
	}
	return phrase
}
