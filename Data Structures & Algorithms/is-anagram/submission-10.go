func isAnagram(s string, t string) bool {

	var a [26]rune

	for _, code := range s {
		a[code - 'a']++
	}

	for _, code := range t {
		a[code - 'a']--
	}

	for _, x := range a {
		if x != 0 {
			return false
		}
	}

	return true
}
