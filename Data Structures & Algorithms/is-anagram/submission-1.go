func isAnagram(s string, t string) bool {
    
    if len(s) != len(t){
        return false
    }

    smap := make(map[rune]int)
	tmap := make(map[rune]int)

	for _, chs := range s {
		smap[chs]++
	}

	for _, cht := range t {
		tmap[cht]++
	}

	for id, rn := range smap{
		if tmap[id] != rn{
			return false
		}
	}
	return true
}