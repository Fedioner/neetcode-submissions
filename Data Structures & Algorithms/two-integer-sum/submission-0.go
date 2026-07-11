func twoSum(nums []int, target int) []int {
    
	mp := make(map[int]int)

	for i, value := range nums {
		if second, ok := mp[target - value]; !ok {
			mp[value] = i
		}else{
			return []int{second, i}
		}
	}
	
	return []int{0,0}
}
