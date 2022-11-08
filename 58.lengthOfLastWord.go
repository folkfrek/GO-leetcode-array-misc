func lengthOfLastWord(s string) int {
    count := 0
    for i:= len(s)-1; i >=0 ; i--{  
        if s[i] == ' ' && count == 0 {
            continue
        } else if s[i] != ' '{
            count++
        } else {
            return count
        }
    }
    return count
}
// Time Complexity: O(n)
// Space Complexity: O(1)
