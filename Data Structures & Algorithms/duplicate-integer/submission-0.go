func hasDuplicate(nums []int) bool {
    value := make(map[int]bool)

    for _, i := range nums {
        if value[i] != false {
            return true
        }else{
            value[i] = true
        }
    }

    return false
}
