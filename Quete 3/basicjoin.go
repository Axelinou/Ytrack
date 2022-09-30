package BasicJoin

func BasicJoin(tab []string) string {
	phrase := ""
	for i := 0; i < len(tab); i++ {
		phrase += tab[i]
	}
	return phrase
}
