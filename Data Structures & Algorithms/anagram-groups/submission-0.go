
func freq(str string) [26]int {
    var result [26]int
    for _, char := range str {
        result[char-'a']++
    }
    return result
}

func groupAnagrams(strs []string) [][]string {

	matrix := make([][]string, 0)
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		groups[freq(str)] = append(groups[freq(str)], str)
	}

	for _, group := range groups {
		matrix = append(matrix, group)
	}

	return matrix
}
