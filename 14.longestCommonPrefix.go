func longestCommonPrefix(strs []string) string {
    str := strs[0]
    for _, char := range strs {
        i, j := 0, 0
        for i < len(str) && j < len(char) {
            if str[i] != char[j] {
                break
            }
            i++
            j++
        }
        str = char[0:j]
    }
    return str
}
// Time Complexity: O(n * k)
// Space Complexity: O(1)
