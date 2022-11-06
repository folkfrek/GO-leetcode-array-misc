func isSubsequence(s string, t string) bool {
    index := 0
    for i := 0; i < len(t); i++ {
        if index < len(s) && s[index] == t[i] {index++}
    }
    if index == len(s) {return true}
    return false
}
// Time Complexity: O(n)
// Space Complexity: O(1)
