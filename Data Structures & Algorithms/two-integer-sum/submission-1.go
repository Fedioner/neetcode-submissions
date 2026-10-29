func twoSum(nums []int, target int) []int {
    
	mp := make(map[int]int)

	for i, num := range nums {
		if first, ok := mp[target - num]; ok {
			return []int{first, i}
		}
		mp[num] = i
	}

	return nil
}
