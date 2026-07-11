
type Element struct {
	Key 	int
	Value 	int
}

func topKFrequent(nums []int, k int) []int {

	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	freqlist := make([]Element, 0, len(freq))

	for key, value := range freq {
		freqlist = append(freqlist, Element{key, value})
	}

	sort.Slice(freqlist, func(i, j int) bool {
        return freqlist[i].Value > freqlist[j].Value
    })

	result := make([]int, k)

	for i := range k {
		result[i] = freqlist[i].Key
	}

	return result
}
