func productExceptSelf(nums []int) []int {
	
	size := len(nums)

	leftPreMul := make([]int, size, cap(nums))
	rightPreMul := make([]int, size, cap(nums))

	leftPreMul[0] = nums[0]
	rightPreMul[size - 1] = nums[size - 1]

	for i, j := 1, size - 2; i < size; i, j = i + 1, j - 1 {
		leftPreMul[i] = nums[i] * leftPreMul[i - 1]
		rightPreMul[j] = nums[j] * rightPreMul[j + 1]
	}

	result := make([]int, size, cap(nums))

	result[0] = rightPreMul[1]
	result[size - 1] = leftPreMul[size - 2]

	for index := 1; index < size - 1; index++ {
		result[index] = leftPreMul[index - 1] * rightPreMul[index + 1]
	}

	return result
}
