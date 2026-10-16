

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteByte('#')
		builder.WriteString(str)
	}
	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1
		result = append(result, encoded[start:start+length])
		i = start + length
	}
	return result
}