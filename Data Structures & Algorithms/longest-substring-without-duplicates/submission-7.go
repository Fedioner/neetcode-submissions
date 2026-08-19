func lengthOfLongestSubstring(s string) int {
    var cnt [256]int

    left := 0
    ans := 0

    for right := 0; right < len(s); right++ {
        cnt[s[right]]++

        for cnt[s[right]] > 1 {
            cnt[s[left]]--
            left++
        }

        ans = max(ans, right-left+1)
    }

    return ans
}