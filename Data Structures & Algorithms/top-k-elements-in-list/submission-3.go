import "slices"
func topKFrequent(nums []int, k int) []int {

	mp := make(map[int]int)

	for _, num := range nums {
		mp[num]++
	}

	type number struct {
		num int
		kol int
	}

	var result []number 
	for num, kol := range mp {
		result = append(result, number{num: num, kol:kol})
	}
	
	slices.SortFunc(result, func(a,b number) int {
		if a.kol == b.kol {
			return 0
		}

		if a.kol < b.kol {
			return 1
		} else {
			return -1
		}
	})

	var res []int
	for _, x := range result[:k] {
		res = append(res, x.num)
	}

	return res

}
