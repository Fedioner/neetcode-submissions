func isPalindrome(s string) bool {
    
    i:= 0
    j:= len(s) - 1

    for i < j {

        for i < j && !IsValid(s[i]){
            i++
        }

        for i < j && !IsValid(s[j]){
            j--
        }

        if toLower(s[i]) != toLower(s[j]) {
            return false
        }

        i++
        j--
    }

    return true
}   

func IsValid(c byte) bool {
    return c >= 'A' && c <= 'Z' ||
        c >= 'a' && c <='z' ||
        c >= '0' && c <= '9' 

}

func toLower(c byte) byte {

    if c >= 'A' && c <= 'Z' {

        return c + ('a' - 'A')

    }

    return c

}
