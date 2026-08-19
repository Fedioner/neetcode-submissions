func search(nums []int, target int) int {

	left, right := 0, len(nums) - 1
	mid := (right - left) / 2 
	
	for left <= right {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
			mid = left + (right - left) / 2 
		} else {
			right = mid - 1
			mid = left + (right - left) / 2 
		} 
	} 

	return -1
}
